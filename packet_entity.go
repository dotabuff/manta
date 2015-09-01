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

// Represents a Packet Entity Event Type
type EntityEventType int

// Possible Packet Entity Event Types
const (
	EntityEventType_Create  = EntityEventType(1)
	EntityEventType_Update  = EntityEventType(2)
	EntityEventType_Destroy = EntityEventType(3)
	EntityEventType_Other   = EntityEventType(4)
)

// Get a property from the entity. Prefers reading from the entity properties,
// falling back to the baseline properties if necessary.
func (pe *PacketEntity) Fetch(key string) (interface{}, bool) {
	v, ok := pe.Properties.Fetch(key)
	if !ok {
		v, ok = pe.ClassBaseline.Fetch(key)
	}

	return v, ok
}

// A function that can handle a game event.
type packetEntityHandler func(*PacketEntity, EntityEventType) error

// Registers a packet entity event handler.
func (p *Parser) OnPacketEntity(fn packetEntityHandler) {
	p.packetEntityHandlers = append(p.packetEntityHandlers, fn)
}

// Internal callback for CSVCMsg_PacketEntities.
func (p *Parser) onCSVCMsg_PacketEntities(m *dota.CSVCMsg_PacketEntities) error {
	// Skip processing if we're configured not to.
	if !p.ProcessPacketEntities {
		return nil
	}

	_debugfl(5, "pTick=%d isDelta=%v deltaFrom=%d updatedEntries=%d maxEntries=%d baseline=%d updateBaseline=%v", p.Tick, m.GetIsDelta(), m.GetDeltaFrom(), m.GetUpdatedEntries(), m.GetMaxEntries(), m.GetBaseline(), m.GetUpdateBaseline())

	// Skip processing full updates after the first. We'll process deltas instead.
	if !m.GetIsDelta() && p.packetEntityFullPackets > 0 {
		return nil
	}

	r := NewReader(m.GetEntityData())
	index := int32(-1)
	pe := &PacketEntity{}
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
		eventType := EntityEventType_Other
		if r.readBoolean() {
			if r.readBoolean() {
				eventType = EntityEventType_Destroy
			} else {
				eventType = EntityEventType_Other
			}
		} else {
			if r.readBoolean() {
				eventType = EntityEventType_Create
			} else {
				eventType = EntityEventType_Update
			}
		}

		_debugfl(5, "update type is %d, %v", eventType, index)

		// Proceed based on the update type
		switch eventType {
		case EntityEventType_Create:
			// Sometimes we're told to create an existing entity.
			// The data doesn't appear to ever change, so just throw it away.
			if pe, ok = p.PacketEntities[index]; ok {
				// We already have an existing entity here, reuse it.
				r.seekBits(p.classIdSize + 25)
			} else {
				// Create a new PacketEntity.
				pe = &PacketEntity{
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
			}

			// Read properties
			pe.Properties.Merge(ReadProperties(r, pe.flatTbl))

		case EntityEventType_Update:
			// Find the existing packetEntity
			pe, ok = p.PacketEntities[index]
			if !ok {
				_panicf("unable to find packet entity %d for update", index)
			}

			// Read properties and update the packetEntity
			pe.Properties.Merge(ReadProperties(r, pe.flatTbl))

		case EntityEventType_Destroy:
			if pe, ok = p.PacketEntities[index]; !ok {
				_panicf("unable to find packet entity %d for delete", index)
			}

			delete(p.PacketEntities, index)
		}

		// Offer the event to callback handlers.
		for _, h := range p.packetEntityHandlers {
			if err := h(pe, eventType); err != nil {
				return err
			}
		}
	}

	// Update the full packet count.
	if !m.GetIsDelta() {
		p.packetEntityFullPackets += 1
	}

	return nil
}
