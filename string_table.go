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
	return nil

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

type stringTable struct {
	name  string
	items []*stringTableItem
}

type stringTableItem struct {
	index int
	key   string
	value []byte
}

const (
	STRINGTABLE_MAX_KEY_SIZE = 1024
	STRINGTABLE_KEY_HISTORY  = 32
)

func parseStringTable(buf []byte, maxEntries int32, userDataFixed bool, userDataSize int32, userDataSizeBits int32) (items []*stringTableItem) {
	items = make([]*stringTableItem, 0)

	// XXX TODO: Clean up after we're done debugging.
	debugMode = false
	defer func() {
		debugMode = true
		if err := recover(); err != nil {
			_debugf("recovered: %s", err)
			return
		}
	}()

	// Create a reader for the buffer
	r := newReader(buf)

	// Start with an index of -1.
	// If the first item is at index 0 it will use a incr operation.
	index := -1

	// Maintain a list of key history
	keys := make([]string, 0, STRINGTABLE_KEY_HISTORY)

	// Some tables have no data
	if len(buf) == 0 {
		return items
	}

	_debugf("index bits = %d", log2(int(maxEntries)))

	// Loop through entries in the data structure
	//
	// Each entry is a tuple consisting of {index, key, value}
	//
	// Index can either be incremented from the previous position or
	// overwritten with a given entry.
	//
	// Key may be omitted (will be represented here as "")
	for i := 0; i < int(maxEntries); i++ {
		key := ""
		value := []byte{}

		// if i > 12 && i < 16 {
		// 	debugMode = true
		// } else {
		// 	debugMode = false
		// }

		// Read a boolean to determine whether the operation is an increment or
		// has a fixed index position. A fixed index position of zero should be
		// the last data in the buffer, and indicates that all data has been read.
		incr := r.readBoolean()
		if incr {
			index++

			_debugf("%d: incr index to %d", i, index)
		} else {
			size := log2(int(maxEntries))
			_debugf("%d: reading %d index bits", i, size) // this might just be 5
			index = int(r.readBits(size))
			_debugf("%d: modify index to %d", i, index)

			// An index of zero given by value indicates the end of the buffer.
			if index == 0 {
				if r.remBits() > 7 {
					_panicf("still have too many (%d) bits left!", r.remBits())
				}

				break
			}
		}

		_debugf("bits left: %d (%d bytes)", r.remBits(), r.remBytes())

		// Some values have keys, some don't.
		hasKey := r.readBoolean()
		if hasKey {
			_debugf("%d: has a key!", i)
			// if full && r.readBoolean() {
			// 	panic("shouldnt happen")
			// }

			substring := r.readBoolean()
			if substring {
				sIndex := r.readBits(5)  // index of substr in keyhistory
				sLength := r.readBits(5) // prefix length to new key

				_debugf("%d: substring index=%d length=%d", i, sIndex, sLength)

				if (sIndex >= STRINGTABLE_KEY_HISTORY) || (sLength >= STRINGTABLE_MAX_KEY_SIZE) {
					panic("shouldnt happen 2.0")
				}

				if int(sIndex) >= len(keys) {
					key += r.readString()
				} else {
					s := keys[sIndex]
					if int(sLength) > len(s) {
						key += s + r.readString()
					} else {
						key += s[0:sLength] + r.readString()
					}
				}
			} else {
				_debugf("%d: reading normal string from byte %d bit %d", i, r.bytePos(), r.pos%8)
				key = r.readString()
			}

			_debugf("%d: key = %s", i, key)

			if len(keys) >= STRINGTABLE_KEY_HISTORY {
				copy(keys[0:], keys[1:])
				keys[len(keys)-1] = ""
				keys = keys[:len(keys)-1]
			}
			keys = append(keys, key)
		} else {
			_debugf("%d: no key", i)
		}

		// read value

		hasValue := r.readBoolean()
		length := 0
		if hasValue {
			_debugf("has value!")

			valSize := 0
			if userDataFixed {
				length = int(userDataSize)
				valSize = int(userDataSizeBits)
				_debugf("fixed length = %d (%d bits) at pos %d (byte %d)", length, valSize, r.pos, r.bytePos())
			} else {
				length = int(r.readBits(14))
				valSize = length * 8
				_debugf("variable length = %d (%d bits) at pos %d (byte %d)", length, valSize, r.pos, r.bytePos())
			}

			value = r.readBitsAsBytes(valSize)
		}

		_debugf("got string %s, len %d", key, len(value))
		items = append(items, &stringTableItem{index, key, value})
	}

	return items
}

func log2(n int) int {
	return int(math.Ceil(math.Log2(float64(n))))
}
