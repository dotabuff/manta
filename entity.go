package manta

import (
	"github.com/dotabuff/manta/dota"
)

const (
	entityCreated = 0x01
	entityUpdated = 0x02
	entityDeleted = 0x04
	entityEntered = 0x08
	entityLeft    = 0x10
)

type entity struct {
	index  int32
	serial int32
	class  *class
	active bool
	state  *fieldState
}

func (p *Parser) onCSVCMsg_PacketEntitiesNew(m *dota.CSVCMsg_PacketEntities) error {
	r := newReader(m.GetEntityData())

	var index = int32(-1)
	var updates = int(m.GetUpdatedEntries())
	var cmd uint32
	var classId int32
	var serial int32

	for ; updates > 0; updates-- {
		index += int32(r.readUBitVar()) + 1

		cmd = r.readBits(2)
		if cmd&0x01 == 0 {
			if cmd&0x02 != 0 {
				classId = int32(r.readBits(p.classIdSize))
				serial = int32(r.readBits(17))
				r.readVarUint32()

				class := p.newClassesById[classId]
				if class == nil {
					_panicf("unable to find new class %d", classId)
				}

				baseline := p.newClassBaselines[classId]
				if baseline == nil {
					_panicf("unable to find new baseline %d", classId)
				}

				entity := &entity{
					index:  index,
					serial: serial,
					class:  class,
					active: true,
					state:  newFieldState(),
				}
				p.newEntities[index] = entity
				readFields(newReader(baseline), class.serializer, entity.state)
				readFields(r, class.serializer, entity.state)

			} else {
				entity := p.newEntities[index]
				if entity == nil {
					_panicf("unable to find existing entity %d", index)
				}
				if !entity.active {
					entity.active = true
				}
				readFields(r, entity.class.serializer, entity.state)
			}

		} else {
			entity := p.newEntities[index]
			if entity != nil {
				if entity.active {
					entity.active = false
				} else {
					_printf("warn: entity %d (%s) ordered to leave, already inactive", entity.class.classId, entity.class.name)
				}

				if cmd&0x02 != 0 {
					p.newEntities[index] = nil
				}

			} else {
				_printf("warn: unable to find existing entity %d", index)
			}
		}
	}

	return nil
}
