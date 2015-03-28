package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

//go:generate go run gen/packet.go dota packet.go

func (p *Parser) OnCDemoPacket(obj *dota.CDemoPacket) error {
	b := NewBitReader(obj.GetData())

	for {
		if b.BytesLeft() <= 1 {
			return nil
		}
		demType, demBytes := b.ReadInnerPacket()

		dem := dota.EDemoCommands(demType)
		if m, err := MessageTypeForEDemoCommands(dem); err == nil {
			if hook, ok := p.hookDEM[dem]; ok {
				callHook(demBytes, m, hook)
			}
		} else {
			PP(dem)
			return (spew.Errorf("unknown message"))
		}
	}
}

func callHook(data []byte, m proto.Message, hook func(proto.Message)) {
	err := proto.Unmarshal(data, m)
	if err != nil {
		panic(err)
	}
	hook(m)
}
