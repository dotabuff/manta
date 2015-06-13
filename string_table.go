package manta

import (
	"github.com/dotabuff/manta/dota"
)

const (
	STRINGTABLE_KEY_HISTORY_SIZE = 32
)

// Holds and maintains the string table information for an
// instance of the Parser.
type stringTables struct {
	tables    map[int32]*stringTable
	nameIndex map[string]int32
	nextIndex int32
}

// Retrieves a string table by its name. Check the bool.
func (ts *stringTables) getTableByName(name string) (*stringTable, bool) {
	i, ok := ts.nameIndex[name]
	if !ok {
		return nil, false
	}
	t, ok := ts.tables[i]
	return t, ok
}

// Creates a new empty stringTables.
func newStringTables() *stringTables {
	return &stringTables{
		tables:    make(map[int32]*stringTable),
		nameIndex: make(map[string]int32),
		nextIndex: 0,
	}
}

// Holds and maintains the information for a string table.
type stringTable struct {
	index             int32
	name              string
	items             map[int32]*stringTableItem
	maxEntries        int32
	userDataFixedSize bool
	userDataSize      int32
}

// Holds and maintains a single entry in a string table.
type stringTableItem struct {
	index int32
	key   string
	value []byte
}

// Internal parser for callback OnCDemoStringTables.
// These appear to be periodic state dumps and appear every 1800 outer ticks.
// XXX TODO: decide if we want to at all integrate these updates,
// or trust create/update entirely. Let's ignore them for now.
func (p *Parser) onCDemoStringTables(m *dota.CDemoStringTables) error {
	return nil
}

// Internal parser for callback OnCSVCMsg_CreateStringTable.
func (p *Parser) onCSVCMsg_CreateStringTable(m *dota.CSVCMsg_CreateStringTable) error {
	// Create a new string table at the next index position
	t := &stringTable{
		index:             p.stringTables.nextIndex,
		name:              m.GetName(),
		items:             make(map[int32]*stringTableItem),
		maxEntries:        m.GetMaxEntries(),
		userDataFixedSize: m.GetUserDataFixedSize(),
		userDataSize:      m.GetUserDataSize(),
	}

	// Increment the index
	p.stringTables.nextIndex += 1

	// Decompress the data if necessary
	buf := m.GetStringData()
	if m.GetDataCompressed() {
		var err error
		if buf, err = unlzss(buf); err != nil {
			return err
		}
	}

	// Parse the items out of the string table data
	items := parseStringTable(buf, m.GetMaxEntries(), t.userDataFixedSize, t.userDataSize)

	// Insert the items into the table
	for _, item := range items {
		t.items[item.index] = item
	}

	// Add the table to the parser state
	p.stringTables.tables[t.index] = t
	p.stringTables.nameIndex[t.name] = t.index

	return nil
}

// Internal parser for callback OnCSVCMsg_UpdateStringTable.
func (p *Parser) onCSVCMsg_UpdateStringTable(m *dota.CSVCMsg_UpdateStringTable) error {
	// TODO: integrate
	t, ok := p.stringTables.tables[m.GetTableId()]
	if !ok {
		_panicf("missing string table %d", m.GetTableId())
	}

	_tracef("tick=%d name=%s changedEntries=%d buflen=%d", p.Tick, t.name, m.GetNumChangedEntries(), len(m.GetStringData()))

	// Parse the updates out of the string table data
	items := parseStringTable(m.GetStringData(), m.GetNumChangedEntries(), t.userDataFixedSize, t.userDataSize)

	// Apply the updates to the parser state
	for _, item := range items {
		if _, ok := t.items[item.index]; ok {
			// XXX TODO: Sometimes ActiveModifiers change keys, which is suspicous...
			if item.key != "" && item.key != t.items[item.index].key {
				_tracef("tick=%d name=%s index=%d key='%s' update key -> %s", p.Tick, t.name, item.index, t.items[item.index].key, item.key)
				t.items[item.index].key = item.key
			}
			if len(item.value) > 0 {
				_tracef("tick=%d name=%s index=%d key='%s' update value len %d -> %d", p.Tick, t.name, item.index, t.items[item.index].key, len(t.items[item.index].value), len(item.value))
				t.items[item.index].value = item.value
			}
		} else {
			_tracef("tick=%d name=%s inserting new item %d key '%s'", p.Tick, t.name, item.index, item.key)
			t.items[item.index] = item
		}
	}

	return nil
}

// Parse a string table data blob, returning a list of item updates.
func parseStringTable(buf []byte, numUpdates int32, userDataFixed bool, userDataSize int32) (items []*stringTableItem) {
	items = make([]*stringTableItem, 0)

	// Create a reader for the buffer
	r := newReader(buf)

	// Start with an index of -1.
	// If the first item is at index 0 it will use a incr operation.
	index := int32(-1)

	// Maintain a list of key history
	keys := make([]string, 0, STRINGTABLE_KEY_HISTORY_SIZE)

	// Some tables have no data
	if len(buf) == 0 {
		return items
	}

	// Loop through entries in the data structure
	//
	// Each entry is a tuple consisting of {index, key, value}
	//
	// Index can either be incremented from the previous position or
	// overwritten with a given entry.
	//
	// Key may be omitted (will be represented here as "")
	//
	// Value may be omitted
	for i := 0; i < int(numUpdates); i++ {
		key := ""
		value := []byte{}

		// Read a boolean to determine whether the operation is an increment or
		// has a fixed index position. A fixed index position of zero should be
		// the last data in the buffer, and indicates that all data has been read.
		incr := r.readBoolean()
		if incr {
			index++
		} else {
			index = int32(r.readVarUint32()) + 1
		}

		// Some values have keys, some don't.
		hasKey := r.readBoolean()
		if hasKey {
			// Some entries use reference a position in the key history for
			// part of the key. If referencing the history, read the position
			// and size from the buffer, then use those to build the string
			// combined with an extra string read (null terminated).
			// Alternatively, just read the string.
			useHistory := r.readBoolean()
			if useHistory {
				pos := r.readBits(5)
				size := r.readBits(5)

				if int(pos) >= len(keys) {
					key += r.readString()
				} else {
					s := keys[pos]
					if int(size) > len(s) {
						key += s + r.readString()
					} else {
						key += s[0:size] + r.readString()
					}
				}
			} else {
				key = r.readString()
			}

			if len(keys) >= STRINGTABLE_KEY_HISTORY_SIZE {
				copy(keys[0:], keys[1:])
				keys[len(keys)-1] = ""
				keys = keys[:len(keys)-1]
			}
			keys = append(keys, key)
		}

		// Some entries have a value.
		hasValue := r.readBoolean()
		if hasValue {
			// Values can be either fixed size (with a size specified in
			// bits during table creation, or have a variable size with
			// a 14-bit prefixed size.
			if userDataFixed {
				value = r.readBitsAsBytes(int(userDataSize))
			} else {
				size := int(r.readBits(14))
				r.readBits(3) // XXX TODO: what is this?
				value = r.readBytes(size)
			}
		}

		items = append(items, &stringTableItem{index, key, value})
	}

	return items
}
