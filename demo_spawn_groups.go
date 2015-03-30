package manta

import "github.com/dotabuff/manta/dota"

func (p *Parser) OnCDemoSpawnGroups(spawnGroups *dota.CDemoSpawnGroups) error {
	b := NewBitReader(spawnGroups.GetData())

	for {
		if b.BytesLeft() <= 1 {
			return nil
		}
		demType, demBytes := b.ReadInnerPacket()

		net := dota.NET_Messages(demType)
		if m, err := MessageTypeForNET_Messages(net); err == nil {
			if hook, ok := p.hookNET[net]; ok {
				callHook(demBytes, m, hook)
			}
		} else {
			PP(net)
			return E("unknown message")
		}
	}
}
