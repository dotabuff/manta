package manta

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

// Internal callback for OnCDemoSendTables.
func (p *Parser) onCDemoSendTables(m *dota.CDemoSendTables) error {
	// Parse the send tables
	st, err := ParseSendTables(m)
	if err != nil {
		return err
	}

	// Update the parser state
	p.SendTables = st

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
	tables map[string]*SendTable
	props  []*SendProp
}

// Get a send table by name.
func (ts *sendTables) GetTableByName(name string) (*SendTable, bool) {
	t, ok := ts.tables[name]
	return t, ok
}

// Represents a single send table.
type SendTable struct {
	Index   int32
	Name    string
	Version int32
	Props   []*SendProp
}

// Represents a property in a send table.
type SendProp struct {
	DtIndex                int32
	DtName                 string
	VarIndex               int32
	VarName                string
	BitCount               *int32
	LowValue               *float32
	HighValue              *float32
	EncodeFlags            *uint32
	FieldSerializerIndex   *int32
	FieldSerializerName    *string
	FieldSerializerVersion *int32
	SerializedFromIndex    *int32
	SerializedFromName     *string
	SendNodeIndex          int32
	SendNodeName           string
}

// Copies a sendProp.
func (p *SendProp) Copy() *SendProp {
	return &SendProp{
		DtIndex:                p.DtIndex,
		DtName:                 p.DtName,
		VarIndex:               p.VarIndex,
		VarName:                p.VarName,
		BitCount:               p.BitCount,
		LowValue:               p.LowValue,
		HighValue:              p.HighValue,
		EncodeFlags:            p.EncodeFlags,
		FieldSerializerIndex:   p.FieldSerializerIndex,
		FieldSerializerName:    p.FieldSerializerName,
		FieldSerializerVersion: p.FieldSerializerVersion,
		SerializedFromIndex:    p.SerializedFromIndex,
		SerializedFromName:     p.SerializedFromName,
		SendNodeIndex:          p.SendNodeIndex,
		SendNodeName:           p.SendNodeName,
	}
}

// Debugging method to describe the table.
func (t *SendTable) Describe() {
	_debugf("table %d (%s) version %d with %d raw props", t.Index, t.Name, t.Version, len(t.Props))
	for i, p := range t.Props {
		_debugf("-> prop %d: %s", i, p.Describe())
	}
}

// Returns the type name and count for a sendprop.
func (p *SendProp) TypeInfo() (s string, n int, err error) {
	ss := strings.Split(strings.Replace(p.DtName, "]", "[", 1), "[")
	s = ss[0]
	switch s {
	case "char":
		n = 1
	default:
		if len(ss) >= 2 {
			n, err = strconv.Atoi(ss[1])
		} else {
			n = 1
		}
	}
	return
}

func (p *SendProp) Describe() string {
	out := _sprintf("type:%s(%d) name:%s(%d)",
		p.DtName, p.DtIndex, p.VarName, p.VarIndex)

	// This doesn't seem to be very helpful yet.
	// out += _sprintf(" sendNode: %s(%d)", p.sendNodeName, p.sendNodeIndex)

	if p.FieldSerializerIndex != nil {
		out += _sprintf(" serializer:%s(%d)", *p.FieldSerializerName, *p.FieldSerializerIndex)
	}

	if p.SerializedFromIndex != nil {
		out += _sprintf(" serializedFrom:%s(%d)", *p.SerializedFromName, *p.SerializedFromIndex)
	}

	if p.BitCount != nil {
		out += _sprintf(" bitCount:%d", *p.BitCount)
	}
	if p.LowValue != nil {
		out += _sprintf(" lowValue:%f", *p.LowValue)
	}
	if p.HighValue != nil {
		out += _sprintf(" highValue:%f", *p.HighValue)
	}
	if p.EncodeFlags != nil {
		out += _sprintf(" encodeFlags:%d", *p.EncodeFlags)
	}
	if p.FieldSerializerVersion != nil {
		out += _sprintf(" fieldSerVer:%d", *p.FieldSerializerVersion)
	}
	return out
}

// Dumps the json representation of the new FlattenedSerializer Packets
func DumpSendTables(m *dota.CDemoSendTables) string {
	// This packet just contains a single large buffer
	r := newReader(m.GetData())

	// The buffer starts with a varint encoded length
	size := int(r.readVarUint32())
	if size != r.remBytes() {
		_panicf("expected %d additional bytes, got %d", size, r.remBytes())
	}

	// Read the rest of the buffer as a CSVCMsg_FlattenedSerializer.
	buf := r.readBytes(size)
	msg := &dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, msg); err != nil {
		_panicf("cannot decode proto: %s", err)
	}

	str, _ := json.MarshalIndent(msg, "", "  ") // two space ident
	return string(str)
}

// Parses a CDemoSendTables buffer, producing a sendTables object.
func ParseSendTables(m *dota.CDemoSendTables) (*sendTables, error) {
	// This packet just contains a single large buffer
	r := newReader(m.GetData())

	// The buffer starts with a varint encoded length
	size := int(r.readVarUint32())
	if size != r.remBytes() {
		_panicf("expected %d additional bytes, got %d", size, r.remBytes())
	}

	// Read the rest of the buffer as a CSVCMsg_FlattenedSerializer.
	buf := r.readBytes(size)
	msg := &dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, msg); err != nil {
		_panicf("cannot decode proto: %s", err)
	}

	// Create a list of sendProps
	props := make([]*SendProp, 0)
	for _, o := range msg.GetFields() {
		p := &SendProp{
			DtIndex:                o.GetVarTypeSym(),
			DtName:                 msg.GetSymbols()[o.GetVarTypeSym()],
			VarIndex:               o.GetVarNameSym(),
			VarName:                msg.GetSymbols()[o.GetVarNameSym()],
			BitCount:               o.BitCount,
			LowValue:               o.LowValue,
			HighValue:              o.HighValue,
			FieldSerializerVersion: o.FieldSerializerVersion,
			SendNodeIndex:          o.GetSendNodeSym(),
			SendNodeName:           msg.GetSymbols()[o.GetSendNodeSym()],
		}

		if o.EncodeFlags != nil {
			p.EncodeFlags = proto.Uint32(uint32(*o.EncodeFlags))
		}

		if o.FieldSerializerNameSym != nil {
			p.FieldSerializerIndex = o.FieldSerializerNameSym
			p.FieldSerializerName = proto.String(msg.GetSymbols()[o.GetFieldSerializerNameSym()])
		}

		props = append(props, p)
	}

	// Create a map of sendTables
	tables := make(map[string]*SendTable)
	for _, o := range msg.GetSerializers() {
		// Create the basic table.
		t := &SendTable{
			Index:   o.GetSerializerNameSym(),
			Name:    msg.GetSymbols()[o.GetSerializerNameSym()],
			Version: o.GetSerializerVersion(),
			Props:   make([]*SendProp, 0),
		}

		// Iterate through prop field indexes.
		for _, pid := range o.GetFieldsIndex() {
			// Get the property at the given index.
			prop := props[int(pid)]

			// If the prop has a serializer, inherit its properties
			if prop.FieldSerializerIndex != nil {
				// Find the serializer.
				ser, ok := tables[*prop.FieldSerializerName]
				if !ok {
					_panicf("unable to find serializer %d (%s)", *prop.FieldSerializerIndex, *prop.FieldSerializerName)
				}

				// Iterate through serializer props, adding them to the table.
				// Property names are subclassed as "%propVarName.%serializerVarName"
				for _, p := range ser.Props {
					p2 := p.Copy()
					p2.SerializedFromIndex = prop.FieldSerializerIndex
					p2.SerializedFromName = prop.FieldSerializerName
					p2.VarName = _sprintf("%s.%s", prop.VarName, p.VarName)
					t.Props = append(t.Props, p2)
				}
				continue
			}

			// For normal props (without serializers), extract base type name
			// and element count, then store those as props in the table.
			typName, typCount, err := prop.TypeInfo()
			if err != nil {
				_panicf("unable to flatten property %s: %s", prop.Describe(), err)
			}

			// Iterate through elements, adding them as new properties to the table.
			// Single element properties are named using the property var name.
			// Multiple element properties are named as "%propVarName.N".
			for i := 0; i < typCount; i++ {
				p2 := prop.Copy()
				p2.DtName = typName
				if typCount > 1 {
					p2.VarName = _sprintf("%s.%04d", prop.VarName, i)
				}
				t.Props = append(t.Props, p2)
			}
		}

		tables[t.Name] = t
	}

	// Return a sendTables object
	return &sendTables{tables: tables, props: props}, nil
}
