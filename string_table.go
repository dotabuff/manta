package manta

import (
	"math"

	"github.com/dotabuff/manta/dota"
)

const (
	STRINGTABLE_KEY_HISTORY_SIZE = 32
)

// Holds and maintains the string table information for an
// instance of the Parser.
type stringTables struct {
	tables    map[string]*stringTable
	nextIndex int32
}

// Creates a new empty stringTables.
func newStringTables() *stringTables {
	return &stringTables{
		tables:    map[string]*stringTable{},
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
func (p *Parser) onCDemoStringTables(m *dota.CDemoStringTables) error {
	// TODO: integrate

	return nil
}

// Internal parser for callback OnCSVCMsg_CreateStringTable.
func (p *Parser) onCSVCMsg_CreateStringTable(m *dota.CSVCMsg_CreateStringTable) error {
	t := &stringTable{
		index:             p.stringTables.nextIndex,
		name:              m.GetName(),
		items:             make(map[int32]*stringTableItem),
		maxEntries:        m.GetMaxEntries(),
		userDataFixedSize: m.GetUserDataFixedSize(),
		userDataSize:      m.GetUserDataSize(),
	}

	p.stringTables.nextIndex += 1

	buf := m.GetStringData()
	if m.GetDataCompressed() {
		var err error
		if buf, err = unlzss(buf); err != nil {
			return err
		}
	}

	items := parseStringTable(buf, t.maxEntries, t.userDataFixedSize, t.userDataSize)
	for _, item := range items {
		t.items[item.index] = item
	}

	return nil
}

// Internal parser for callback OnCSVCMsg_UpdateStringTable.
func (p *Parser) onCSVCMsg_UpdateStringTable(m *dota.CSVCMsg_UpdateStringTable) error {
	// TODO: integrate

	return nil
}

func parseStringTable(buf []byte, maxEntries int32, userDataFixed bool, userDataSize int32) (items []*stringTableItem) {
	items = make([]*stringTableItem, 0)

	// Create a reader for the buffer
	r := newReader(buf)

	// Start with an index of -1.
	// If the first item is at index 0 it will use a incr operation.
	index := int32(-1)
	indexBits := log2(int(maxEntries))

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
	for i := 0; i < int(maxEntries); i++ {
		key := ""
		value := []byte{}

		// Read a boolean to determine whether the operation is an increment or
		// has a fixed index position. A fixed index position of zero should be
		// the last data in the buffer, and indicates that all data has been read.
		incr := r.readBoolean()
		if incr {
			index++
		} else {
			// XXX TODO: This path is untested and should not be trusted.
			index = int32(r.readBits(indexBits))
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

func log2(n int) int {
	return int(math.Ceil(math.Log2(float64(n))))
}
