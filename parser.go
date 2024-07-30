package manta

import (
	"bytes"
	"fmt"
	"io"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/snappy"
)

// The first 8 bytes of a replay for Source 1 and Source 2
var magicSource1 = []byte{'P', 'U', 'F', 'D', 'E', 'M', 'S', '\000'}
var magicSource2 = []byte{'P', 'B', 'D', 'E', 'M', 'S', '2', '\000'}

// Parser is an instance of the replay parser
type Parser struct {
	// Callbacks provide a mechanism for receiving notification
	// when a specific message has been received and decoded.
	Callbacks *Callbacks

	// Contains the game tick associated with the last message processed.
	Tick uint32

	// Contains the net tick associated with the last net message processed.
	NetTick uint32

	// Stores the game build.
	GameBuild uint32

	// AfterStopCallback is a function to be called when the parser stops.
	AfterStopCallback func()

	classBaselines             map[int32][]byte
	classesById                map[int32]*class
	classesByName              map[string]*class
	classIdSize                uint32
	classInfo                  bool
	entities                   map[int32]*Entity
	entityFullPackets          int
	entityHandlers             []EntityHandler
	gameEventHandlers          map[string][]GameEventHandler
	gameEventNames             map[int32]string
	gameEventTypes             map[string]*gameEventType
	isStopping                 bool
	modifierTableEntryHandlers []ModifierTableEntryHandler
	serializers                map[string]*serializer
	stream                     *stream
	stringTables               *stringTables
	stopAtTick                 uint32
}

// Create a new parser from a byte slice.
func NewParser(buf []byte) (*Parser, error) {
	r := bytes.NewReader(buf)
	return NewStreamParser(r)
}

// Create a new Parser from an io.Reader
func NewStreamParser(r io.Reader) (*Parser, error) {
	// Create a new parser with an internal reader for the given buffer.
	parser := &Parser{
		Callbacks: newCallbacks(),
		Tick:      0,
		NetTick:   0,
		GameBuild: 0,

		classBaselines:    make(map[int32][]byte),
		classesById:       make(map[int32]*class),
		classesByName:     make(map[string]*class),
		entities:          make(map[int32]*Entity),
		entityHandlers:    make([]EntityHandler, 0),
		gameEventHandlers: make(map[string][]GameEventHandler),
		gameEventNames:    make(map[int32]string),
		gameEventTypes:    make(map[string]*gameEventType),
		isStopping:        false,
		serializers:       make(map[string]*serializer),
		stream:            newStream(r),
		stringTables:      newStringTables(),
	}

	// Parse out the header, ensuring that it's valid.
	magic, err := parser.stream.readBytes(8)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(magic, magicSource2) {
		return nil, _errorf("unexpected magic: expected %s, got %s", magicSource2, magic)
	}

	// Skip the next 8 bytes, which appear to be two int32s related to the size
	// of the demo file. We may need them in the future, but not so far.
	parser.stream.readBytes(8)

	// Internal handlers
	parser.Callbacks.OnCDemoPacket(parser.onCDemoPacket)
	parser.Callbacks.OnCDemoSignonPacket(parser.onCDemoPacket)
	parser.Callbacks.OnCDemoFullPacket(parser.onCDemoFullPacket)
	parser.Callbacks.OnCSVCMsg_CreateStringTable(parser.onCSVCMsg_CreateStringTable)
	parser.Callbacks.OnCSVCMsg_UpdateStringTable(parser.onCSVCMsg_UpdateStringTable)
	parser.Callbacks.OnCSVCMsg_ServerInfo(parser.onCSVCMsg_ServerInfo)
	parser.Callbacks.OnCMsgSource1LegacyGameEventList(parser.onCMsgSource1LegacyGameEventList)
	parser.Callbacks.OnCMsgSource1LegacyGameEvent(parser.onCMsgSource1LegacyGameEvent)

	parser.Callbacks.OnCDemoClassInfo(parser.onCDemoClassInfo)
	parser.Callbacks.OnCDemoSendTables(parser.onCDemoSendTables)
	parser.Callbacks.OnCSVCMsg_PacketEntities(parser.onCSVCMsg_PacketEntities)

	// Maintains the value of parser.Tick
	parser.Callbacks.OnCNETMsg_Tick(func(m *dota.CNETMsg_Tick) error {
		parser.NetTick = m.GetTick()
		return nil
	})

	return parser, nil
}

// Start parsing the replay. Will stop processing new events after Stop() is called.
func (p *Parser) Start() (err error) {
	var msg *outerMessage

	defer p.afterStop()

	defer func() {
		if p := recover(); p != nil {
			if e, ok := p.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", p)
			}
		}
	}()

	// Loop through all outer messages until we're signaled to stop. Stopping
	// happens when either the OnCDemoStop message is encountered or
	// parser.Stop() is called programatically.
	for !p.isStopping {
		if p.stopAtTick > 0 && p.Tick > p.stopAtTick {
			return
		}

		msg, err = p.readOuterMessage()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}

		p.Tick = msg.tick

		if err = p.Callbacks.callByDemoType(msg.typeId, msg.data); err != nil {
			return
		}
	}

	return
}

// Stop parsing the replay, causing the parser to stop processing new events.
func (p *Parser) Stop() {
	p.isStopping = true
}

func (p *Parser) afterStop() {
	if p.AfterStopCallback != nil {
		p.AfterStopCallback()
	}
}

// Performs a lookup on a string table by an entry index.
func (p *Parser) LookupStringByIndex(table string, index int32) (string, bool) {
	t, ok := p.stringTables.GetTableByName(table)
	if !ok {
		return "", false
	}

	item, ok := t.Items[index]
	if !ok {
		return "", false
	}

	return item.Key, true
}

// Describes a demo message parsed from the replay.
type outerMessage struct {
	tick   uint32
	typeId int32
	data   []byte
}

// Read the next outer message from the buffer.
func (p *Parser) readOuterMessage() (*outerMessage, error) {
	// Read a command header, which includes both the message type
	// well as a flag to determine whether or not whether or not the
	// message is compressed with snappy.
	command, err := p.stream.readCommand()
	if err != nil {
		return nil, err
	}

	// Extract the type and compressed flag out of the command
	msgType := int32(command & ^dota.EDemoCommands_DEM_IsCompressed)
	msgCompressed := (command & dota.EDemoCommands_DEM_IsCompressed) == dota.EDemoCommands_DEM_IsCompressed

	// Read the tick that the message corresponds with.
	tick, err := p.stream.readVarUint32()
	if err != nil {
		return nil, err
	}

	// This appears to actually be an int32, where a -1 means pre-game.
	if tick == 4294967295 {
		tick = 0
	}

	// Read the size and following buffer.
	size, err := p.stream.readVarUint32()
	if err != nil {
		return nil, err
	}

	buf, err := p.stream.readBytes(size)
	if err != nil {
		return nil, err
	}

	// If the buffer is compressed, decompress it with snappy.
	if msgCompressed {
		var err error
		if buf, err = snappy.Decode(nil, buf); err != nil {
			return nil, err
		}
	}

	// Return the message
	msg := &outerMessage{
		tick:   tick,
		typeId: msgType,
		data:   buf,
	}
	return msg, nil
}

// parseToTick configures this Parser to stop once it has parsed the given tick.
func (p *Parser) parseToTick(n uint32) {
	p.stopAtTick = n
}
