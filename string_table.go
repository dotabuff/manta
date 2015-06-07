package manta

import (
	"github.com/dotabuff/manta/dota"
)

// Internal parser for callback OnCDemoStringTables.
func (p *Parser) onCDemoStringTables(m *dota.CDemoStringTables) error {
	for _, t := range m.GetTables() {
		switch t.GetTableName() {
		case "userinfo":
			if err := stParseUserInfo(t); err != nil {
				return err
			}
		}
	}

	return nil
}

// Internal parser for callback OnCSVCMsg_CreateStringTable.
func (p *Parser) onCSVCMsg_CreateStringTable(m *dota.CSVCMsg_CreateStringTable) error {
	return nil
}

// Internal parser for callback OnCSVCMsg_UpdateStringTable.
func (p *Parser) onCSVCMsg_UpdateStringTable(m *dota.CSVCMsg_UpdateStringTable) error {
	return nil
}

// TODO XXX: WIP
// So far:
// - First byte (calling it flags) reads in as a int8 or varint, value 10
// - Second byte (calling it length) reads in as a int8 or varint and provides the length of the following string
// - Next item is a string with the player name
// - The byte after the name reads in as int8 or varint, value 17
func stParseUserInfo(t *dota.CDemoStringTablesTableT) error {
	_debugf("stParseUserInfo flags %d", t.GetTableFlags())
	for i, item := range t.GetItems() {
		if len(item.GetData()) == 0 {
			continue
		}

		r := newReader(item.GetData())
		flags := r.read_var_uint64()
		nameLen := r.read_var_uint64()
		name := r.read_string_n(int(nameLen))
		_dump("user info buffer", item.GetData())
		_debugf("index=%d flags=%d name_len=%d name='%s'", i, flags, nameLen, name)
	}

	return nil
}
