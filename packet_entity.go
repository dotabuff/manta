package manta

import (
	"github.com/dotabuff/manta/dota"
)

const (
	peUpdateTypeCreate   = 1
	peUpdateTypePreserve = 2
	peUpdateTypeDelete   = 3
	peUpdateTypeLeave    = 4
)

type PacketEntity struct {
	Tick         uint32
	Index        int
	SerialNum    int
	ClassId      int32
	ClassName    string
	EntityHandle int
	Name         string
	UpdateType   int
	Values       map[string]interface{}
}

// Internal parser for callback CSVCMsg_PacketEntities.
func (p *Parser) onCSVCMsg_PacketEntities(m *dota.CSVCMsg_PacketEntities) error {
	// XXX TODO: continue exploring
	return nil

	r := newPacketEntityReader(p, m.GetEntityData())
	for i := 0; i < int(m.GetUpdatedEntries()); i++ {
		r.readNextPacketEntity()
	}

	return nil
}

type packetEntityReader struct {
	p     *Parser
	r     *reader
	index int
}

// Creates a new packetEntityReader, used to read data in the expected format.
func newPacketEntityReader(p *Parser, buf []byte) *packetEntityReader {
	return &packetEntityReader{p, newReader(buf), -1}
}

// Reads the index for the next packet entity, updating the internal reader state.
// Status: this appears to work for the first entity. Don't trust it.
func (r *packetEntityReader) readNextIndex() {
	x := r.r.read_bits(4)
	b1 := r.r.read_boolean()
	b2 := r.r.read_boolean()

	if b1 {
		x += (r.r.read_bits(4) << 4)
	}
	if b2 {
		x += (r.r.read_bits(8) << 4)
	}

	r.index += 1 + int(x)
}

// Determines the PE update type by reading two booleans
// Status: this appears to work for the first entity. Don't trust it.
func (r *packetEntityReader) readUpdateType() int {
	b1 := r.r.read_boolean()
	b2 := r.r.read_boolean()

	switch {
	case b1 && b2:
		return peUpdateTypeDelete
	case b1 && !b2:
		return peUpdateTypeLeave
	case !b1 && b2:
		return peUpdateTypeCreate
	case !b1 && !b2:
		return peUpdateTypePreserve
	}

	// impossible
	return -1
}

// Reads the next packet entity from the buffer
func (r *packetEntityReader) readNextPacketEntity() *PacketEntity {
	// Each entity in the buffer is prefixed by an indexing command,
	// read that first and update the internal reader state.
	r.readNextIndex()
	_debugf("pe index %d", r.index)

	// Each entity in the buffer can be one of many update types, which
	// dictates the encoding and application. Get the update type here.
	updateType := r.readUpdateType()
	_debugf("pe update type %d", updateType)

	var pe *PacketEntity

	switch updateType {
	case peUpdateTypeCreate:
		pe = &PacketEntity{
			Tick:       r.p.Tick,
			ClassId:    int32(r.r.read_bits(r.p.classIdSize)), // assumption from yasha, 10.
			SerialNum:  int(r.r.read_bits(10)),                // assumption from yasha
			Index:      r.index,
			UpdateType: updateType,
			Values:     make(map[string]interface{}),
		}
		pe.ClassName = r.p.classInfo[pe.ClassId]

	case peUpdateTypePreserve:
		// XXX TODO: this should look up an existing PE from the Parser
		pe = &PacketEntity{
			Tick:       r.p.Tick,
			UpdateType: updateType,
		}
	case peUpdateTypeDelete:
		// XXX TODO
	case peUpdateTypeLeave:
		// XXX TODO
	}

	if updateType == peUpdateTypeCreate || updateType == peUpdateTypePreserve {
		// XXX TODO: do something with this after we get it reading properly.
		r.readPropertiesIndex()
	}

	return pe
}

// Reads the property index for the current PE from the buffer.
// Status: this is a direct port from source1 logic and is not working,
// or reading at the wrong position. Don't trust it.
func (r *packetEntityReader) readPropertiesIndex() []int {
	props := []int{}
	prop := -1

	for {
		if r.r.read_boolean() {
			prop += 1
			props = append(props, prop)
			_debugf("easy %d", prop)
		} else {
			x := r.r.read_var_uint32()
			if x == 16383 {
				break
			}
			prop += 1 + int(x)
			props = append(props, prop)
			_debugf("hard %d", prop)
		}
	}

	return props
}
