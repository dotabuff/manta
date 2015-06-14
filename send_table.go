package manta

import (
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

// Internal callback for OnCDemoSendTables.
func (p *Parser) onCDemoSendTables(m *dota.CDemoSendTables) error {
	// Parse the send tables
	st, err := parseSendTables(m)
	if err != nil {
		return err
	}

	// Update the parser state
	p.sendTables = st

	return nil
}

// Internal callback for OnCSVCMsg_SendTable.
// So far we haven't seen any of these, they very well may not exist.
func (p *Parser) onCSVCMsg_SendTable(m *dota.CSVCMsg_SendTable) error {
	_panicf("got a CSVCMsg_SendTable")
	return nil
}

// Holds and maintains send tables for an instance of Parser.
type sendTables struct {
	tables map[string]*sendTable
	props  []*sendProp
}

// Get a send table by name.
func (ts *sendTables) getTableByName(name string) (*sendTable, bool) {
	t, ok := ts.tables[name]
	return t, ok
}

// Represents a single send table.
type sendTable struct {
	index    int32
	name     string
	unknown2 int32
	props    []*sendProp
}

// Represents a property in a send table.
type sendProp struct {
	dtIndex   int32
	dtName    string
	varIndex  int32
	varName   string
	lowValue  *float32
	highValue *float32
	unknown3  *int32
	unknown6  *int32
	unknown7  *int32
	unknown8  *int32
	unknown9  int32
}

// Parses a CDemoSendTables buffer, producing a sendTables object.
func parseSendTables(m *dota.CDemoSendTables) (*sendTables, error) {
	// This packet just contains a single large buffer
	r := newReader(m.GetData())

	// The buffer starts with a varint encoded length
	size := int(r.readVarUint32())
	if size != r.remBytes() {
		_panicf("expected %d additional bytes, got %d", size, r.remBytes())
	}

	// The rest of the buffer contains a repeated structure where each entry
	// contains a type (varint), size (varint) and buffer of the given size.
	//
	// So far we've encountered three types of packets:
	//
	// - Type 10 appears to be a SendTable. It has a string index, an unknown
	//   field (probably flags), and a repeating list of property indexes.
	//   This is encoded as a protobuf message.
	//
	// - Type 18 is a string, which appears to be indexed by the nature order
	//   of items in the message.
	//
	// - Type 26 appears to be a SendProp. It has 9 fields, most of which are
	//   optional. This is encoded as a protobuf message. Field descriptions:
	//
	//   - The first field (a varint) contains the index of a string that
	//     represents the type (ex. "Vector", "bool", "int32", "MoveCollide_t",
	//     "CHandle< CBaseCombatWeapon >", etc...)
	//
	//   - The second field (also a varint) contains the index of a string that
	//     represents the name of the field (ex. "m_vecSpecifiedSurroundingMaxs",
	//     "m_bSimulatedEveryTick", "m_MoveCollide", "m_fEffects", etc...)
	//
	//   - The third field appears to be bit size, but we're not sure quite yet.
	//
	//   - The fourth and fifth fields appear to be lowValue and highValue.
	//
	//   - The 6th, 7th, and 8th fields are usually not present. We're not yet
	//     sure what they are used for.
	//
	//   - The 9th field (a varint) contains the index of a string that **MIGHT**
	//     represent a base type, or relation of some kind. Rationale:
	//
	//     - All fields have a 9th field
	//     - All entries in the 9th field correspond to an indexed string.
	//     - Most values for the 9th field are grouped by type (ex. most int32's
	//       share the same value for field 9).
	//
	//     It's also possible that the 9th field contains bitmasked flags.
	//

	wireTables := make([]*wireSendTable, 0)
	wireProps := make([]*wireSendProp, 0)
	wireStrings := make([]string, 0)

	for i := 0; r.remBytes() > 0; i++ {
		// Read the type
		t := r.readVarUint32()

		// Then the length-prefixed buffer
		size := int(r.readVarUint32())
		buf := r.readBytes(size)

		// Handle the message based on type
		switch t {
		case 10:
			// Type 10 appears to be a SendTable-like thing.
			o := &wireSendTable{}
			if err := proto.Unmarshal(buf, o); err != nil {
				_debugf("unable to unmarshal a thing10: %s", err)
				continue
			}
			wireTables = append(wireTables, o)

		case 18:
			// Type 18 is an indexed string.
			wireStrings = append(wireStrings, string(buf))

		case 26:
			// Type 26 appears to be a SendProp-like thing.
			o := &wireSendProp{}
			if err := proto.Unmarshal(buf, o); err != nil {
				_debugf("unable to unmarshal a thing10: %s", err)
				continue
			}
			wireProps = append(wireProps, o)

		default:
			_panicf("unknown type %d", t)
		}
	}

	// Create a list of sendProps
	props := make([]*sendProp, 0)
	for _, o := range wireProps {
		p := &sendProp{
			dtIndex:   o.GetTypeIndex(),
			dtName:    wireStrings[int(o.GetTypeIndex())],
			varIndex:  o.GetVarIndex(),
			varName:   wireStrings[int(o.GetVarIndex())],
			lowValue:  o.LowValue,
			highValue: o.HighValue,
			unknown3:  o.Unknown3,
			unknown6:  o.Unknown6,
			unknown7:  o.Unknown7,
			unknown8:  o.Unknown8,
			unknown9:  o.GetUnknown9(),
		}
		props = append(props, p)
	}

	// Create a list of sendTables
	tables := make(map[string]*sendTable)
	for _, o := range wireTables {
		t := &sendTable{
			index:    o.GetIndex(),
			name:     wireStrings[int(o.GetIndex())],
			unknown2: o.GetUnknown2(),
			props:    make([]*sendProp, len(o.GetProperties())),
		}
		for i, pid := range o.GetProperties() {
			t.props[i] = props[int(pid)]
		}
		tables[t.name] = t
	}

	// Return a sendTables object
	return &sendTables{tables: tables, props: props}, nil
}
