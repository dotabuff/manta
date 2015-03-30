package manta

import "github.com/dotabuff/manta/dota"

func (p *Parser) OnCDemoStop(stop *dota.CDemoStop) { p.Stop() }
func (p *Parser) OnCDemoFileHeader(fileHeader *dota.CDemoFileHeader) {
	PP(fileHeader)
}

func (p *Parser) OnCDemoSendTables(sendTables *dota.CDemoSendTables) {
	b := NewBitReader(sendTables.GetData())
	l := b.ReadVarInt()
	PP(l)

	for {
		if b.BytesLeft() <= 1 {
			return
		}
		demType, demBytes := b.ReadInnerPacket()

		PP(demType, demBytes)
	}
}

/*
	buf := proto.NewBuffer(sendTables.GetData())
	buf.DecodeVarint()

	for {
		raw, err := buf.DecodeRawBytes(false)
		if err != nil {
			break
		}

		m := &dota.CDemoSendTables{}
		err = proto.Unmarshal(raw, m)
		if err != nil {
			panic(err)
		}
		PP(m)
	}
	/*
		b := proto.NewBuffer(sendTables.GetData())
		for {
			iLen, err := b.DecodeVarint()
			if err != nil {
				panic(err)
			}
			PP(iLen)
			iTyp, err := b.DecodeVarint()
			if err != nil {
				panic(err)
			}
			PP(iTyp)

			m := &dota.CSVCMsg_SendTable{}
			err = b.Unmarshal(m)
			if err != nil {
				panic(err)
			}

			return
		}
*/
