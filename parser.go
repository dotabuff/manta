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
	fd, err := os.Open("/home/manveru/Dropbox/Dotabuff/dotabuff_rocks.dem")
	if err != nil {
		panic(err)
	}

	parser := &Parser{
		sendTableId: -1,
		stream:      bufio.NewReader(fd),
	}

	if err = parser.ReadHeader(); err != nil {
		panic(err)
	}

	for !parser.isStopping {
		if msg, err := parser.Read(); err != nil {
			panic(err)
		} else {
			if err = parser.OnRawMessage(msg); err != nil {
				panic(err)
			}
		}
	}
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

func (p *Parser) ReadHeader() error {
	header := DemoHeader{}
	err := binary.Read(p.stream, binary.LittleEndian, &header)
	if err != nil {
		return err
	}

	if header.Magic != ProtobufDemoSource2Magic {
		panic(spew.Sprintf("Expected %v but got %v", ProtobufDemoSource2Magic, header.Magic))
	}
	p.Header = header

	return nil
}

func (p *Parser) Read() (Message, error) {
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

func (p *Parser) OnRawMessage(msg Message) error {
	switch msg.Type {
	case dota.EDemoCommands_DEM_FileHeader:
		return p.OnProtoMessage(msg, &dota.CDemoFileHeader{})
	case dota.EDemoCommands_DEM_SignonPacket:
		return p.OnProtoMessage(msg, &dota.CDemoPacket{})
	case dota.EDemoCommands_DEM_SendTables:
		return p.OnProtoMessage(msg, &dota.CDemoSendTables{})
	case dota.EDemoCommands_DEM_ClassInfo:
		return p.OnProtoMessage(msg, &dota.CDemoClassInfo{})
	case dota.EDemoCommands_DEM_ConsoleCmd:
		return p.OnProtoMessage(msg, &dota.CDemoConsoleCmd{})
	case dota.EDemoCommands_DEM_FileInfo:
		return p.OnProtoMessage(msg, &dota.CDemoFileInfo{})
	case dota.EDemoCommands_DEM_Packet:
		return p.OnProtoMessage(msg, &dota.CDemoPacket{})
	case dota.EDemoCommands_DEM_SpawnGroups:
		return p.OnProtoMessage(msg, &dota.CDemoSpawnGroups{})
	case dota.EDemoCommands_DEM_Stop:
		return p.OnProtoMessage(msg, &dota.CDemoStop{})
	case dota.EDemoCommands_DEM_StringTables:
		return p.OnProtoMessage(msg, &dota.CDemoStringTables{})
	case dota.EDemoCommands_DEM_SyncTick:
		return p.OnProtoMessage(msg, &dota.CDemoSyncTick{})
	case dota.EDemoCommands_DEM_UserCmd:
		return p.OnProtoMessage(msg, &dota.CDemoUserCmd{})
	}
	return spew.Errorf("Unhandled Raw Message: %v", msg)
}

type HasGetData interface {
	GetData() []byte
}

func (p *Parser) OnProtoMessage(msg Message, obj proto.Message) error {
	if err := proto.Unmarshal(msg.data, obj); err != nil {
		return err
	}
	msg.data = nil

	/*
		isHasGetData, ok := obj.(HasGetData)
		if ok {
			PP(msg.Compressed, msg.Type, isHasGetData.GetData())
		} else {
			PP(msg.Compressed, msg.Type, obj)
		}
	*/

	switch typedObject := obj.(type) {
	case *dota.CDemoFileHeader:
		p.OnCDemoFileHeader(msg, typedObject)
	case *dota.CDemoPacket:
		p.OnCDemoPacket(msg, typedObject)
	case *dota.CDemoSendTables:
		p.OnCDemoSendTables(msg, typedObject)
	case *dota.CDemoClassInfo:
		p.OnCDemoClassInfo(msg, typedObject)
	case *dota.CDemoSyncTick:
		p.OnCDemoSyncTick(msg, typedObject)
	case *dota.CDemoStringTables:
		p.OnCDemoStringTables(msg, typedObject)
	case *dota.CDemoSpawnGroups:
		p.OnCDemoSpawnGroups(msg, typedObject)
	case *dota.CDemoConsoleCmd:
		p.OnCDemoConsoleCmd(msg, typedObject)
	case *dota.CDemoUserCmd:
		p.OnCDemoUserCmd(msg, typedObject)
	case *dota.CDemoStop:
		p.OnCDemoStop(msg, typedObject)
	default:
		PP(msg, obj)
	}

	return nil
}

var (
	P = spew.Dump
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
