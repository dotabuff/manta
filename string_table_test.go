package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestParseStringTable(t *testing.T) {
	assert := assert.New(t)

	scenarios := []struct {
		fixturePath string
		tableName   string
		itemCount   int
		firstItem   *stringTableItem
		lastItem    *stringTableItem
	}{
		// PASSING: CombatLogNames is uncompressed and has 24 entries (working)
		//{
		//	"string_tables_17_335_uncompressed.pbmsg",
		//	"CombatLogNames",
		//	24,
		//	&stringTableItem{0, "dota_unknown", []byte{}},
		//	&stringTableItem{23, "item_flask", []byte{}},
		//},

		// PASSING: downloadables is uncompressed and has no entries (working)
		// {
		// 	"string_tables_1_29_uncompressed.pbmsg",
		// 	"downloadables",
		// 	0,
		// 	nil,
		// 	nil,
		// },

		// PASSING: ResponseKeys is uncompressed and has 15 entries (working)
		// {
		// 	"string_tables_18_175_uncompressed.pbmsg",
		// 	"ResponseKeys",
		// 	15,
		// 	&stringTableItem{0, "concept", []byte{}},
		// 	&stringTableItem{14, "game_start_time", []byte{}},
		// },

		// PASSING: server_query_info is uncompressed and has 1 entry
		// {
		// 	"string_tables_7_50_uncompressed.pbmsg",
		// 	"server_query_info",
		// 	1,
		// 	&stringTableItem{0, "QueryPort", []byte{0x0, 0x0, 0x0, 0x0}},
		// 	&stringTableItem{0, "QueryPort", []byte{0x0, 0x0, 0x0, 0x0}},
		// },

		// FAILING: lightstyles is compressed and has NNNNNNN entries
		// {
		// 	"string_tables_5_590_compressed.pbmsg",
		// 	"lightstyles",
		// 	1,
		// 	&stringTableItem{0, "", []byte{}},
		// 	&stringTableItem{0, "", []byte{}},
		// },

		// FAILING: instancebaseline is compressed and has NNNNNNN entries
		// {
		// 	"string_tables_4_22356_compressed.pbmsg",
		// 	"instancebaseline",
		// 	1,
		// 	&stringTableItem{0, "", []byte{}},
		// 	&stringTableItem{0, "", []byte{}},
		// },

		// FAILING: EntityNames is compressed
		{
			"string_tables_8_4162_compressed.pbmsg",
			"EntityNames",
			0,
			&stringTableItem{0, "", []byte{}},
			&stringTableItem{0, "", []byte{}},
		},

		// FAILING: EconItems is not compressed and fails on values
		// {
		// 	"string_tables_16_559_uncompressed.pbmsg",
		// 	"EconItems",
		// 	0,
		// 	nil,
		// 	nil,
		// },
	}

	// Iterate through test scenarios
	for _, s := range scenarios {
		// Load the message from the fixture
		m := &dota.CSVCMsg_CreateStringTable{}
		err := proto.Unmarshal(_read_fixture(s.fixturePath), m)
		if err != nil {
			t.Errorf("unable to decode %s: %s", s.fixturePath, err)
			continue
		}

		// Decompress the data if need be
		buf := m.GetStringData()
		if m.GetDataCompressed() {
			buf, err = decompress2(buf)
			if err != nil {
				t.Errorf("unable to decompress %s: %s", s.fixturePath, err)
				continue
			}
		}

		// Make sure we're looking at the right table
		assert.Equal(s.tableName, m.GetName())

		// XXX TODO: remove
		_debugf("buflen=%d max_entries=%d fixed=%d size=%d size_bits=%d",
			len(buf), m.GetMaxEntries(), m.GetUserDataFixedSize(), m.GetUserDataSize(), m.GetUserDataSizeBits())

		// Parse the table data
		items := parseStringTable(buf, m.GetMaxEntries(), m.GetUserDataFixedSize() == 1, m.GetUserDataSize(), m.GetUserDataSizeBits())

		// Make sure we have the correct number of entries
		assert.Equal(s.itemCount, len(items))

		// Verify the first and last entries
		if s.firstItem != nil {
			assert.Equal(s.firstItem, items[0])
		}
		if s.lastItem != nil {
			assert.Equal(s.lastItem, items[len(items)-1])
		}
	}
}
