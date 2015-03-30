package manta

import "github.com/dotabuff/manta/dota"

func (p *Parser) onCDemoSpawnGroups(spawnGroups *dota.CDemoSpawnGroups) error {
	return nil
	b := NewBitReader(spawnGroups.GetData())

	for {
		if b.BytesLeft() <= 1 {
			return nil
		}
		demType, demBytes := b.ReadInnerPacket()

		PP(demType, demBytes)
		/*
			net := dota.NET_Messages(demType)
			if m, err := MessageTypeForNET_Messages(net); err == nil {
				if hook, ok := p.hookNET[net]; ok {
					callHook(demBytes, m, hook)
				}
			} else {
				PP(net)
				return E("unknown message")
			}
		*/
	}
}
