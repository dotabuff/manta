package manta

import (
	"os"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

var DEBUG bool

func init() {
	if os.Getenv("DEBUG") != "" {
		DEBUG = true
	}
}

//go:generate go run gen/packet.go dota message_lookup.go

func (p *Parser) onCDemoPacket(obj *dota.CDemoPacket) error {
	b := NewBitReader(obj.GetData())

	for {
		if b.BytesLeft() <= 1 {
			return nil
		}
		demType, demBytes := b.ReadInnerPacket()

		p.CallByPacketType(demType, demBytes)
	}
}

func callHook(data []byte, m proto.Message, hook func(proto.Message)) {
	err := proto.Unmarshal(data, m)
	if err != nil {
		panic(err)
	}
	hook(m)
}
