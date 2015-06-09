package manta

import (
	"github.com/dotabuff/manta/dota"
)

// Internal parser for callback OnCDemoSendTables.
func (p *Parser) onCDemoSendTables(m *dota.CDemoSendTables) error {
	_debugf("got a CDemoSendTables")
	return parseSendTables(m)
}

// Internal parser for callback OnCSVCMsg_SendTable
func (p *Parser) onCSVCMsg_SendTable(m *dota.CSVCMsg_SendTable) error {
	_debugf("got a CSVCMsg_SendTable")
	// XXX TODO: manage props
	return nil
}

func parseSendTables(m *dota.CDemoSendTables) error {
	// This packet just contains a single large buffer
	r := newReader(m.GetData())

	// The buffer starts with a varint encoded length
	size := int(r.read_var_uint32())
	if size != r.rem_bytes() {
		_panicf("expected %d additional bytes, got %d", size, r.rem_bytes())
	}

	// XXX TODO:
	// The rest of the structure is only semi-known.
	// By reading a varint type, varint size and fixed sized object we
	// get some useful data, some junk data, but it all appears somewhat well
	// aligned. The type identifiers don't seem to line up to anything very
	// meaningful - maybe we need new protos, or maybe we're reading the struct
	// wrong.
	// This technique produces about 5000 short messages.
	// The reads align: we don't under-read or over-read the buffer.
	for r.rem_bytes() > 0 {
		r.read_var_uint32()           // type
		s := int(r.read_var_uint32()) // length
		r.read_bytes(s)               // buffer
	}

	return nil
}
