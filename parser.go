package manta

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
)

func NewParserFromFile(path string) *Parser {
	fd, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	parser := &Parser{
		sendTableId:  -1,
		stream:       bufio.NewReader(fd),
		classInfo:    map[int]string{},
		stringTables: NewStringTables(),
		Callbacks:    &Callbacks{},
	}

	cb := parser.Callbacks

	cb.OnCNETMsg_SpawnGroup_Load = func(m *dota.CNETMsg_SpawnGroup_Load) error { parser.onSpawnGroupLoad(m); return nil }
	cb.OnCSVCMsg_UpdateStringTable = func(m *dota.CSVCMsg_UpdateStringTable) error { parser.stringTables.onUpdateStringTable(m); return nil }
	cb.OnSignonPacket = func(m *dota.CDemoPacket) error { parser.onCDemoPacket(m); return nil }
	cb.OnCDemoClassInfo = func(m *dota.CDemoClassInfo) error { parser.onCDemoClassInfo(m); return nil }
	cb.OnCDemoStop = func(m *dota.CDemoStop) error { parser.Stop(); return nil }
	cb.OnCDemoPacket = func(m *dota.CDemoPacket) error { parser.onCDemoPacket(m); return nil }
	cb.OnCDemoSpawnGroups = func(m *dota.CDemoSpawnGroups) error { parser.onCDemoSpawnGroups(m); return nil }
	cb.OnCDemoStringTables = func(m *dota.CDemoStringTables) error { parser.stringTables.onCDemoStringTables(m); return nil }
	cb.OnCDemoUserCmd = func(m *dota.CDemoUserCmd) error { parser.onCDemoUserCmd(m); return nil }

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
	stream       *bufio.Reader
	Header       DemoHeader
	isStopping   bool
	sendTableId  int64
	classInfo    map[int]string
	stringTables *StringTables
	Callbacks    *Callbacks
}

func (p *Parser) Start() {
	for !p.isStopping {
		if msg, err := p.read(); err != nil {
			panic(err)
		} else {
			if err = p.CallByDemoType(int32(msg.Type), msg.data); err != nil {
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
