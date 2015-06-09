package manta

import (
	"math"

	"github.com/dotabuff/manta/dota"
)

func NewStringTables() *StringTables {
	return &StringTables{
		tables: map[string]StringTable{},
	}
}

type StringTables struct {
	tables map[string]StringTable
}

type StringTable struct {
	items map[int]*dota.CDemoStringTablesItemsT
}

func (p *StringTables) onCDemoStringTables(stringTables *dota.CDemoStringTables) error {
	tables := map[string]StringTable{}

	for _, st := range stringTables.GetTables() {
		table := StringTable{items: map[int]*dota.CDemoStringTablesItemsT{}}
		tables[st.GetTableName()] = table
		for n, item := range st.GetItems() {
			table.items[n] = item
		}
	}

	return nil
}

// Internal parser for callback OnCDemoStringTables.
func (p *Parser) onCDemoStringTables(m *dota.CDemoStringTables) error {
	for _, t := range m.GetTables() {
		switch t.GetTableName() {
		default:
			_debugf("ignoring string table %s with flags %d, %d/%d items", t.GetTableName(), t.GetTableFlags(), len(t.GetItems()), len(t.GetItemsClientside()))
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

type StringTableItem struct {
	Name string
	Data []byte
}

// Parse the entries in a string table
func parseStringTable(m *dota.CSVCMsg_CreateStringTable) map[int]*StringTableItem {
	items := make(map[int]*StringTableItem)

	buf := m.GetStringData()
	if len(buf) == 0 {
		return items
	}

	r := newReader(buf)

	// Some values are compressed and include a header containing LZSS and a uint32
	// denoting the length of the entire file.
	if m.GetDataCompressed() {
		if h1 := r.read_string_n(4); h1 != "LZSS" {
			_panicf("expected LZSS header, got %s", h1)
		}

		if h2 := r.read_le_uint32(); int(h2) != len(buf) {
			_panicf("expected %d length header, got %d", len(buf), h2)
		}
	}

	// This is all guesswork ported from Yasha.

	dataFixedSize := m.GetUserDataFixedSize() == 1
	dataSizeBits := m.GetUserDataSizeBits()

	bitsPerIndex := int(math.Log(float64(m.GetMaxEntries())) / math.Log(2))
	keyHistory := make([]string, 0, 32)
	mysteryFlag := r.read_boolean()
	index := -1

	for r.rem_bits() > 3 {
		if r.read_boolean() {
			index += 1
		} else {
			index = int(r.read_bits(bitsPerIndex))
		}

		name := ""

		if r.read_boolean() {

			if mysteryFlag && r.read_boolean() {
				panic("mysteryFlag assertion failed!")
			}

			if r.read_boolean() {

				basis := r.read_bits(5)
				length := r.read_bits(5)
				if int(basis) >= len(keyHistory) {

					name += r.read_string()
				} else {

					s := keyHistory[basis]
					if int(length) > len(s) {
						name += s + r.read_string()

					} else {
						name += s[0:length] + r.read_string()

					}
				}
			} else {
				name += r.read_string()
			}

			if len(keyHistory) >= 32 {
				copy(keyHistory[0:], keyHistory[1:])
				keyHistory[len(keyHistory)-1] = ""
				keyHistory = keyHistory[:len(keyHistory)-1]
			}
			keyHistory = append(keyHistory, name)
		}

		value := []byte{}
		if r.read_boolean() {
			bitLen := 0
			if dataFixedSize {
				bitLen = int(dataSizeBits)

			} else {
				bitLen = int(r.read_bits(14) * 8)

			}
			value = r.read_bits_as_bytes(bitLen)
		}
		items[index] = &StringTableItem{Name: name, Data: value}
	}

	return items
}
