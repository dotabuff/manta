package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
	"runtime"
	"strings"

	"code.google.com/p/snappy-go/snappy"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

func main() {
	for _, arg := range os.Args[1:] {
		parser := NewParserFromFile(arg)
		parser.HookNET(dota.NET_Messages_net_Tick, func(m proto.Message) { PP(m) })
		parser.Start()
	}
}

func (p *Parser) HookNET(mType dota.NET_Messages, callback func(m proto.Message)) {
	p.hookNET[mType] = callback
}

func NewParserFromFile(path string) *Parser {
	fd, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	parser := &Parser{
		sendTableId: -1,
		stream:      bufio.NewReader(fd),
	}

	parser.hookNET = map[dota.NET_Messages]func(proto.Message){
		dota.NET_Messages_net_SpawnGroup_Load: func(m proto.Message) { parser.OnSpawnGroupLoad(m.(*dota.CNETMsg_SpawnGroup_Load)) },
	}
	parser.hookSVC = map[dota.SVC_Messages]func(proto.Message){}
	parser.hookDUM = map[dota.EDotaUserMessages]func(proto.Message){}
	parser.hookBEM = map[dota.EBaseEntityMessages]func(proto.Message){}
	parser.hookBUM = map[dota.EBaseUserMessages]func(proto.Message){}
	parser.hookBGE = map[dota.EBaseGameEvents]func(proto.Message){}
	parser.hookDEM = map[dota.EDemoCommands]func(proto.Message){
		dota.EDemoCommands_DEM_FileHeader:   func(m proto.Message) {},
		dota.EDemoCommands_DEM_SignonPacket: func(m proto.Message) {},
		dota.EDemoCommands_DEM_SendTables:   func(m proto.Message) {},
		dota.EDemoCommands_DEM_ClassInfo:    func(m proto.Message) {},
		dota.EDemoCommands_DEM_ConsoleCmd:   func(m proto.Message) {},
		dota.EDemoCommands_DEM_FileInfo:     func(m proto.Message) { parser.Stop() },
		dota.EDemoCommands_DEM_Packet:       func(m proto.Message) {},
		dota.EDemoCommands_DEM_SpawnGroups:  func(m proto.Message) { parser.OnCDemoSpawnGroups(m.(*dota.CDemoSpawnGroups)) },
		dota.EDemoCommands_DEM_Stop:         func(m proto.Message) {},
		dota.EDemoCommands_DEM_StringTables: func(m proto.Message) {},
		dota.EDemoCommands_DEM_SyncTick:     func(m proto.Message) {},
		dota.EDemoCommands_DEM_UserCmd:      func(m proto.Message) { parser.OnCDemoUserCmd(m.(*dota.CDemoUserCmd)) },
	}

	if err = parser.readHeader(); err != nil {
		panic(err)
	}

	return parser
}

type DemoHeader struct {
	Magic         [8]byte
	SummaryOffset int32
}

type Parser struct {
	stream      *bufio.Reader
	Header      DemoHeader
	isStopping  bool
	sendTableId int64
	hookDEM     map[dota.EDemoCommands]func(proto.Message)
	hookNET     map[dota.NET_Messages]func(proto.Message)
	hookSVC     map[dota.SVC_Messages]func(proto.Message)
	hookDUM     map[dota.EDotaUserMessages]func(proto.Message)
	hookBEM     map[dota.EBaseEntityMessages]func(proto.Message)
	hookBUM     map[dota.EBaseUserMessages]func(proto.Message)
	hookBGE     map[dota.EBaseGameEvents]func(proto.Message)
}

func (p *Parser) Start() {
	for !p.isStopping {
		if msg, err := p.read(); err != nil {
			panic(err)
		} else {
			if err = p.HandleDemoMessage(msg); err != nil {
				panic(err)
			}
		}
	}
}

func (p *Parser) Stop() {
	p.isStopping = true
}

type Message struct {
	Compressed bool
	Tick       uint64
	Type       dota.EDemoCommands
	data       []byte
	Size       uint64
}

var ProtobufDemoSource2Magic = [8]byte{'P', 'B', 'D', 'E', 'M', 'S', '2', '\000'}

func (p *Parser) readHeader() error {
	header := DemoHeader{}
	err := binary.Read(p.stream, binary.LittleEndian, &header)
	if err != nil {
		return err
	}

	if header.Magic != ProtobufDemoSource2Magic {
		panic(spew.Sprintf("Expected %s but got %s", ProtobufDemoSource2Magic, header.Magic))
	}
	p.Header = header

	return nil
}

func (p *Parser) read() (Message, error) {
	msg := Message{}

	binType, err := binary.ReadUvarint(p.stream)
	if err != nil {
		return msg, err
	}

	binTick, err := binary.ReadUvarint(p.stream)
	if err != nil {
		return msg, err
	}
	msg.Tick = binTick

	binSize, err := binary.ReadUvarint(p.stream)
	if err != nil {
		return msg, err
	}
	msg.Size = binSize

	command := dota.EDemoCommands(binType)
	msg.Compressed = (command & dota.EDemoCommands_DEM_IsCompressed) == dota.EDemoCommands_DEM_IsCompressed
	msg.Type = command & ^dota.EDemoCommands_DEM_IsCompressed

	if binSize > 0x100000 {
		return msg, spew.Errorf("buffer too big: %d", binSize)
	}

	buf := make([]byte, binSize)
	readLen, err := io.ReadFull(p.stream, buf)
	if err != nil {
		return msg, err
	}
	if uint64(readLen) != binSize {
		return msg, spew.Errorf("readLen %d != binSize %d", readLen, binSize)
	}

	if msg.Compressed {
		decodedLen, err := snappy.DecodedLen(buf)
		if err != nil {
			return msg, err
		}

		if decodedLen > 0x100000 {
			return msg, spew.Errorf("decompressed size too big: %d", decodedLen)
		}

		out, err := snappy.Decode(nil, buf)
		if err != nil {
			return msg, err
		}
		msg.data = out
		msg.Size = uint64(decodedLen)
	} else {
		msg.data = buf
	}

	return msg, nil
}

func (p *Parser) HandleDemoMessage(msg Message) error {
	if _, ok := dota.EDemoCommands_name[int32(msg.Type)]; ok {
		if hook, ok := p.hookDEM[msg.Type]; ok {
			m, err := MessageTypeForEDemoCommands(msg.Type)
			if err != nil {
				return err
			}
			if err := proto.Unmarshal(msg.data, m); err != nil {
				return err
			}
			hook(m)
			return nil
		} else {
			return spew.Errorf("Unhandled Raw Message: %v", msg)
		}
	}
	return spew.Errorf("Unknown Raw Message: %v", msg)
}

var (
	P = spew.Dump
	E = spew.Errorf
)

// use this to make debugging output that shows the location.
func PP(args ...interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		f := runtime.FuncForPC(pc)
		fParts := strings.Split(f.Name(), ".")
		fun := fParts[len(fParts)-1]
		s := spew.Sprintf("vvvvvvvvvvvvvvv %s vvvvvvvvvvvvvvv\n", fun)
		spew.Print(s)
		spew.Dump(args...)
		spew.Println(strings.Repeat("^", len(s)-1))
	} else {
		spew.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
		spew.Dump(args...)
		spew.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	}
}
