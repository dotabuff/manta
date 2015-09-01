package manta

import (
	"github.com/dotabuff/manta/dota"
)

// Represents the state of an entity
type PacketEntity struct {
	Index         int32
	ClassId       int32
	ClassName     string
	ClassBaseline *Properties
	Properties    *Properties
	Serial        int32

	flatTbl *dt
}

// Get a property from the entity. Prefers reading from the entity properties,
// falling back to the baseline properties if necessary.
func (pe *PacketEntity) Fetch(key string) (interface{}, bool) {
	v, ok := pe.Properties.Fetch(key)
	if !ok {
		v, ok = pe.ClassBaseline.Fetch(key)
	}

	return v, ok
}

// Internal callback for CSVCMsg_PacketEntities.
func (p *Parser) onCSVCMsg_PacketEntities(m *dota.CSVCMsg_PacketEntities) error {
	if !p.ProcessPacketEntities {
		return nil
	}

	_debugfl(5, "pTick=%d isDelta=%v deltaFrom=%d updatedEntries=%d maxEntries=%d baseline=%d updateBaseline=%v", p.Tick, m.GetIsDelta(), m.GetDeltaFrom(), m.GetUpdatedEntries(), m.GetMaxEntries(), m.GetBaseline(), m.GetUpdateBaseline())

	r := NewReader(m.GetEntityData())
	index := int32(-1)
	ok := false

	// Iterate over all entries
	for i := 0; i < int(m.GetUpdatedEntries()); i++ {
		// Read the index delta from the buffer. This is an implementation
		// from Alice. An alternate implementation from Yasha has the same result.
		delta := r.readUBitVar()
		index += int32(delta) + 1
		_debugfl(5, "index delta is %d to %d", delta, index)

		// Read the type of update based on two booleans.
		// This appears to be backwards from source 1:
		// true+true used to be "create", now appears to be false+true?
		// This seems suspcious.
		updateType := " "
		if r.readBoolean() {
			if r.readBoolean() {
				updateType = "D"
			} else {
				updateType = "?"
			}
		} else {
			if r.readBoolean() {
				updateType = "C"
			} else {
				updateType = "U"
			}
		}

		_debugfl(5, "update type is %s, %v", updateType, index)

		// Proceed based on the update type
		switch updateType {
		case "C":
			// Create a new PacketEntity.
			pe := &PacketEntity{
				Index:      index,
				ClassId:    int32(r.readBits(p.classIdSize)),
				Serial:     int32(r.readBits(25)),
				Properties: NewProperties(),
			}

			// Get the associated class
			if pe.ClassName, ok = p.ClassInfo[pe.ClassId]; !ok {
				_panicf("unable to find class %d", pe.ClassId)
			}

			// Get the associated baseline
			if pe.ClassBaseline, ok = p.ClassBaselines[pe.ClassId]; !ok {
				_panicf("unable to find class baseline %d", pe.ClassId)
			}

			// Get the associated serializer
			if pe.flatTbl, ok = p.serializers[pe.ClassName][0]; !ok {
				_panicf("unable to find serializer for class %s", pe.ClassName)
			}

			// Register the packetEntity with the parser.
			p.PacketEntities[index] = pe

			// Read properties
			pe.Properties = ReadProperties(r, pe.flatTbl)

		case "U":
			// Find the existing packetEntity
			pe, ok := p.PacketEntities[index]
			if !ok {
				_panicf("unable to find packet entity %d for update", index)
			}

			// Read properties and update the packetEntity
			pe.Properties.Merge(ReadProperties(r, pe.flatTbl))

		case "D":
			if _, ok := p.PacketEntities[index]; !ok {
				_panicf("unable to find packet entity %d for delete", index)
			}

			delete(p.PacketEntities, index)
		}
	}

	return nil
}
