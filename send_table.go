package manta

import (
	"strconv"
	"strings"

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
	index   int32
	name    string
	version int32
	props   []*sendProp
}

// Represents a property in a send table.
type sendProp struct {
	dtIndex                int32
	dtName                 string
	varIndex               int32
	varName                string
	bitCount               *int32
	lowValue               *float32
	highValue              *float32
	encodeFlags            *int32
	fieldSerializerIndex   *int32
	fieldSerializerName    *string
	fieldSerializerVersion *int32
	sendNodeIndex          int32
	sendNodeName           string
}

// Returns the type name and count for a sendprop.
func (p *sendProp) typeInfo() (s string, n int, err error) {
	ss := strings.Split(strings.Replace(p.dtName, "]", "[", 1), "[")
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

func (p *sendProp) Describe() string {
	out := _sprintf("type:%s(%d) name:%s(%d)",
		p.dtName, p.dtIndex, p.varName, p.varIndex)

	// This doesn't seem to be very helpful yet.
	// out += _sprintf(" sendNode: %s(%d)", p.sendNodeName, p.sendNodeIndex)

	if p.fieldSerializerIndex != nil {
		out += _sprintf(" serializer:%s(%d)", *p.fieldSerializerName, *p.fieldSerializerIndex)
	}

	if p.bitCount != nil {
		out += _sprintf(" bitCount:%d", *p.bitCount)
	}
	if p.lowValue != nil {
		out += _sprintf(" lowValue:%f", *p.lowValue)
	}
	if p.highValue != nil {
		out += _sprintf(" highValue:%f", *p.highValue)
	}
	if p.encodeFlags != nil {
		out += _sprintf(" encodeFlags:%d", *p.encodeFlags)
	}
	if p.fieldSerializerVersion != nil {
		out += _sprintf(" fieldSerVer:%d", *p.fieldSerializerVersion)
	}
	return out
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

	// Read the rest of the buffer as a CSVCMsg_FlattenedSerializer.
	buf := r.readBytes(size)
	msg := &dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, msg); err != nil {
		_panicf("cannot decode proto: %s", err)
	}

	// Create a list of sendProps
	props := make([]*sendProp, 0)
	for _, o := range msg.GetFields() {
		p := &sendProp{
			dtIndex:                o.GetVarTypeSym(),
			dtName:                 msg.GetSymbols()[o.GetVarTypeSym()],
			varIndex:               o.GetVarNameSym(),
			varName:                msg.GetSymbols()[o.GetVarNameSym()],
			bitCount:               o.BitCount,
			lowValue:               o.LowValue,
			highValue:              o.HighValue,
			encodeFlags:            o.EncodeFlags,
			fieldSerializerVersion: o.FieldSerializerVersion,
			sendNodeIndex:          o.GetSendNodeSym(),
			sendNodeName:           msg.GetSymbols()[o.GetSendNodeSym()],
		}

		if o.FieldSerializerNameSym != nil {
			p.fieldSerializerIndex = o.FieldSerializerNameSym
			p.fieldSerializerName = proto.String(msg.GetSymbols()[o.GetFieldSerializerNameSym()])
		}

		props = append(props, p)
	}

	// Create a map of sendTables
	tables := make(map[string]*sendTable)
	for _, o := range msg.GetSerializers() {
		t := &sendTable{
			index:   o.GetSerializerNameSym(),
			name:    msg.GetSymbols()[o.GetSerializerNameSym()],
			version: o.GetSerializerVersion(),
			props:   make([]*sendProp, len(o.GetFieldsIndex())),
		}
		for i, pid := range o.GetFieldsIndex() {
			t.props[i] = props[int(pid)]
		}
		tables[t.name] = t
	}

	// Return a sendTables object
	return &sendTables{tables: tables, props: props}, nil
}
