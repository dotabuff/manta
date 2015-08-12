package manta

import (
	"bytes"
	"io/ioutil"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/snappy"
)

// The first 8 bytes of a replay for Source 1 and Source 2
var magicSource1 = []byte{'P', 'U', 'F', 'D', 'E', 'M', 'S', '\000'}
var magicSource2 = []byte{'P', 'B', 'D', 'E', 'M', 'S', '2', '\000'}

// A replay parser capable of parsing Source 2 replays
type Parser struct {
	// Callbacks provide a mechanism for receiving notification
	// when a specific message has been received and decoded.
	Callbacks *Callbacks

	// Contains the game tick associated with the last message processed.
	Tick uint32

	// Contains the net tick associated with the last net message processed.
	NetTick uint32

	hasClassInfo      bool
	ClassInfo         map[int32]string
	classIdSize       int
	ClassBaseline     map[int32]map[string]interface{}
	packetEntities    map[int32]*packetEntity
	SendTables        *SendTables
	StringTables      *StringTables
	Serializers       map[string]map[int32]*dt
	spawnGroups       map[uint32]*spawnGroup
	gameEventNames    map[int32]string
	gameEventTypes    map[string]*gameEventType
	gameEventHandlers map[string][]gameEventHandler

	reader            *Reader
	isStopping        bool
	AfterStopCallback func()
}

// Create a new Parser from a file on disk.
func NewParserFromFile(path string) (*Parser, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewParser(buf)
}

// Create a new parser from a byte slice.
func NewParser(buf []byte) (*Parser, error) {
	// Create a new parser with an internal reader for the given buffer.
	parser := &Parser{
		Callbacks: &Callbacks{},

		Tick:    0,
		NetTick: 0,

		reader:     NewReader(buf),
		isStopping: false,

		ClassInfo:         make(map[int32]string),
		ClassBaseline:     make(map[int32]map[string]interface{}),
		packetEntities:    make(map[int32]*packetEntity),
		StringTables:      newStringTables(),
		spawnGroups:       make(map[uint32]*spawnGroup),
		gameEventNames:    make(map[int32]string),
		gameEventTypes:    make(map[string]*gameEventType),
		gameEventHandlers: make(map[string][]gameEventHandler),
	}

	// Parse out the header, ensuring that it's valid.
	if magic := parser.reader.readBytes(8); !bytes.Equal(magic, magicSource2) {
		return nil, _errorf("unexpected magic: expected %s, got %s", magicSource2, magic)
	}

	// Skip the next 8 bytes, which appear to be two int32s related to the size
	// of the demo file. We may need them in the future, but not so far.
	parser.reader.seekBytes(8)

	// Internal handlers
	parser.Callbacks.OnCDemoPacket(parser.onCDemoPacket)
	parser.Callbacks.OnCDemoSignonPacket(parser.onCDemoPacket)
	parser.Callbacks.OnCDemoFullPacket(parser.onCDemoFullPacket)
	parser.Callbacks.OnCDemoClassInfo(parser.onCDemoClassInfo)
	parser.Callbacks.OnCDemoSendTables(parser.onCDemoSendTablesNew)
	parser.Callbacks.OnCSVCMsg_CreateStringTable(parser.onCSVCMsg_CreateStringTable)
	parser.Callbacks.OnCSVCMsg_PacketEntities(parser.onCSVCMsg_PacketEntities)
	parser.Callbacks.OnCSVCMsg_SendTable(parser.onCSVCMsg_SendTable)
	parser.Callbacks.OnCSVCMsg_UpdateStringTable(parser.onCSVCMsg_UpdateStringTable)
	parser.Callbacks.OnCSVCMsg_ServerInfo(parser.onCSVCMsg_ServerInfo)
	parser.Callbacks.OnCNETMsg_SpawnGroup_Load(parser.onCNETMsg_SpawnGroup_Load)
	parser.Callbacks.OnCNETMsg_SpawnGroup_ManifestUpdate(parser.onCNETMsg_SpawnGroup_ManifestUpdate)
	parser.Callbacks.OnCNETMsg_SpawnGroup_SetCreationTick(parser.onCNETMsg_SpawnGroup_SetCreationTick)
	parser.Callbacks.OnCNETMsg_SpawnGroup_Unload(parser.onCNETMsg_SpawnGroup_Unload)
	parser.Callbacks.OnCNETMsg_SpawnGroup_LoadCompleted(parser.onCNETMsg_SpawnGroup_LoadCompleted)
	parser.Callbacks.OnCMsgSource1LegacyGameEventList(parser.onCMsgSource1LegacyGameEventList)
	parser.Callbacks.OnCMsgSource1LegacyGameEvent(parser.onCMsgSource1LegacyGameEvent)

	// Panic if we see any of these
	parser.Callbacks.OnCSVCMsg_GameEvent(func(m *dota.CSVCMsg_GameEvent) error {
		_panicf("unexpected: saw a CSVCMsg_GameEvent")
		return nil
	})

	// Maintains the value of parser.Tick
	parser.Callbacks.OnCNETMsg_Tick(func(m *dota.CNETMsg_Tick) error {
		parser.NetTick = m.GetTick()
		return nil
	})

	// Stops parsing when we reach the end of the replay.
	parser.Callbacks.OnCDemoStop(func(m *dota.CDemoStop) error {
		parser.Stop()
		return nil
	})

	return parser, nil
}

// Start parsing the replay. Will stop processing new events after Stop() is called.
func (p *Parser) Start() error {
	var msg *outerMessage
	var err error

	defer p.afterStop()

	// Loop through all outer messages until we're signaled to stop. Stopping
	// happens when either the OnCDemoStop message is encountered or
	// parser.Stop() is called programatically.
	for !p.isStopping {
		// Read the next outer message.
		if msg, err = p.readOuterMessage(); err != nil {
			return err
		}

		// Update the parser tick
		p.Tick = msg.tick

		// Invoke callbacks for the given message type.
		if err = p.CallByDemoType(msg.typeId, msg.data); err != nil {
			return err
		}
	}

	return nil
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
	t, ok := p.StringTables.GetTableByName(table)
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
	command := dota.EDemoCommands(p.reader.readVarUint32())

	// Extract the type and compressed flag out of the command
	msgType := int32(command & ^dota.EDemoCommands_DEM_IsCompressed)
	msgCompressed := (command & dota.EDemoCommands_DEM_IsCompressed) == dota.EDemoCommands_DEM_IsCompressed

	// Read the tick that the message corresponds with.
	tick := p.reader.readVarUint32()

	// This appears to actually be an int32, where a -1 means pre-game.
	if tick == 4294967295 {
		tick = 0
	}

	// Read the size and following buffer.
	size := int(p.reader.readVarUint32())
	buf := p.reader.readBytes(size)

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
