package manta

import (
	"github.com/dotabuff/manta/dota"
	"github.com/golang/snappy"
)

const (
	stringtableKeyHistorySize = 32
)

// Holds and maintains the string table information for an
// instance of the Parser.
type stringTables struct {
	Tables    map[int32]*stringTable
	NameIndex map[string]int32
	nextIndex int32
}

// Retrieves a string table by its name. Check the bool.
func (ts *stringTables) GetTableByName(name string) (*stringTable, bool) {
	i, ok := ts.NameIndex[name]
	if !ok {
		return nil, false
	}
	t, ok := ts.Tables[i]
	return t, ok
}

// Creates a new empty StringTables.
func newStringTables() *stringTables {
	return &stringTables{
		Tables:    make(map[int32]*stringTable),
		NameIndex: make(map[string]int32),
		nextIndex: 0,
	}
}

// Holds and maintains the information for a string table.
type stringTable struct {
	index             int32
	name              string
	Items             map[int32]*stringTableItem
	userDataFixedSize bool
	userDataSize      int32
	flags             int32
	varintBitCounts   bool
}

func (st *stringTable) GetIndex() int32                      { return st.index }
func (st *stringTable) GetName() string                      { return st.name }
func (st *stringTable) GetItem(index int32) *stringTableItem { return st.Items[index] }

// Holds and maintains a single entry in a string table.
type stringTableItem struct {
	Index int32
	Key   string
	Value []byte
}

// Internal callback for CDemoStringTables.
// These appear to be periodic state dumps and appear every 1800 outer ticks.
// XXX TODO: decide if we want to at all integrate these updates,
// or trust create/update entirely. Let's ignore them for now.
func (p *Parser) onCDemoStringTables(m *dota.CDemoStringTables) error {
	return nil
}

// Internal callback for CSVCMsg_CreateStringTable.
// XXX TODO: This is currently using an artificial, internally crafted message.
// This should be replaced with the real message once we have updated protos.
func (p *Parser) onCSVCMsg_CreateStringTable(m *dota.CSVCMsg_CreateStringTable) error {
	// Create a new string table at the next index position
	t := &stringTable{
		index:             p.stringTables.nextIndex,
		name:              m.GetName(),
		Items:             make(map[int32]*stringTableItem),
		userDataFixedSize: m.GetUserDataFixedSize(),
		userDataSize:      m.GetUserDataSize(),
		flags:             m.GetFlags(),
		varintBitCounts:   m.GetUsingVarintBitcounts(),
	}

	// Increment the index
	p.stringTables.nextIndex += 1

	// Decompress the data if necessary
	buf := m.GetStringData()
	if m.GetDataCompressed() {
		// old replays = lzss
		// new replays = snappy

		r := newReader(buf)
		var err error

		if s := r.readStringN(4); s != "LZSS" {
			if buf, err = snappy.Decode(nil, buf); err != nil {
				return err
			}
		} else {
			if buf, err = unlzss(buf); err != nil {
				return err
			}
		}
	}

	// Parse the items out of the string table data
	items := parseStringTable(buf, m.GetNumEntries(), t.name, t.userDataFixedSize, t.userDataSize, t.flags, t.varintBitCounts)

	// Insert the items into the table
	for _, item := range items {
		t.Items[item.Index] = item
	}

	// Add the table to the parser state
	p.stringTables.Tables[t.index] = t
	p.stringTables.NameIndex[t.name] = t.index

	// Apply the updates to baseline state
	if t.name == "instancebaseline" {
		p.updateInstanceBaseline()
	}

	// Emit events for modifier table entry updates
	if t.name == "ActiveModifiers" {
		if err := p.emitModifierTableEvents(items); err != nil {
			return err
		}
	}

	return nil
}

// Internal callback for CSVCMsg_UpdateStringTable.
func (p *Parser) onCSVCMsg_UpdateStringTable(m *dota.CSVCMsg_UpdateStringTable) error {
	// TODO: integrate
	t, ok := p.stringTables.Tables[m.GetTableId()]
	if !ok {
		_panicf("missing string table %d", m.GetTableId())
	}

	if v(5) {
		_debugf("tick=%d name=%s changedEntries=%d size=%d", p.Tick, t.name, m.GetNumChangedEntries(), len(m.GetStringData()))
	}

	// Parse the updates out of the string table data
	items := parseStringTable(m.GetStringData(), m.GetNumChangedEntries(), t.name, t.userDataFixedSize, t.userDataSize, t.flags, t.varintBitCounts)

	// Apply the updates to the parser state
	for _, item := range items {
		index := item.Index
		if _, ok := t.Items[index]; ok {
			if item.Key != "" && item.Key != t.Items[index].Key {
				t.Items[index].Key = item.Key
			}
			if len(item.Value) > 0 {
				t.Items[index].Value = item.Value
			}
		} else {
			t.Items[index] = item
		}
	}

	// Apply the updates to baseline state
	if t.name == "instancebaseline" {
		p.updateInstanceBaseline()
	}

	// Emit events for modifier table entry updates
	if t.name == "ActiveModifiers" {
		if err := p.emitModifierTableEvents(items); err != nil {
			return err
		}
	}

	return nil
}

// Parse a string table data blob, returning a list of item updates.
func parseStringTable(buf []byte, numUpdates int32, name string, userDataFixed bool, userDataSize int32, flags int32, varintBitCounts bool) (items []*stringTableItem) {
	defer func() {
		if err := recover(); err != nil {
			_debugf("warning: unable to parse string table %s: %s", name, err)
			return
		}
	}()

	items = make([]*stringTableItem, 0)

	// Create a reader for the buffer
	r := newReader(buf)

	// Start with an index of -1.
	// If the first item is at index 0 it will use a incr operation.
	index := int32(-1)

	// Maintain a list of key history
	keys := make([]string, 0, stringtableKeyHistorySize)

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

			if len(keys) >= stringtableKeyHistorySize {
				copy(keys[0:], keys[1:])
				keys[len(keys)-1] = ""
				keys = keys[:len(keys)-1]
			}
			keys = append(keys, key)
		}

		// Some entries have a value.
		hasValue := r.readBoolean()
		if hasValue {
			bitSize := uint32(0)
			isCompressed := false
			if userDataFixed {
				bitSize = uint32(userDataSize)
			} else {
				if (flags & 0x1) != 0 {
					isCompressed = r.readBoolean()
				}
				if varintBitCounts {
					bitSize = r.readUBitVar() * 8
				} else {
					bitSize = r.readBits(17) * 8
				}
			}
			value = r.readBitsAsBytes(bitSize)

			if isCompressed {
				tmp, err := snappy.Decode(nil, value)
				if err != nil {
					_panicf("unable to decode snappy compressed stringtable item (%s, %d, %s): %s", name, index, key, err)
				}
				value = tmp
			}
		}

		items = append(items, &stringTableItem{index, key, value})
	}

	return items
}
