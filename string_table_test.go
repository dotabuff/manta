package manta

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/imprint-esports/manta/dota"
	"github.com/stretchr/testify/assert"
)

func TestParseStringTableCreate(t *testing.T) {
	// XXX: string tables changed, skipping for now.
	t.Skip()

	assert := assert.New(t)

	scenarios := []struct {
		fixturePath string
		tableName   string
		itemCount   int
		firstItem   *stringTableItem
		lastItem    *stringTableItem
	}{
		// CombatLogNames is uncompressed and has 24 entries (working)
		{
			"17_335_uncompressed.pbmsg",
			"CombatLogNames",
			24,
			&stringTableItem{0, "dota_unknown", []byte{}},
			&stringTableItem{23, "item_flask", []byte{}},
		},

		// downloadables is uncompressed and has no entries (working)
		{
			"01_29_uncompressed.pbmsg",
			"downloadables",
			0,
			nil,
			nil,
		},

		// ResponseKeys is uncompressed and has 15 entries (working)
		{
			"18_175_uncompressed.pbmsg",
			"ResponseKeys",
			15,
			&stringTableItem{0, "concept", []byte{}},
			&stringTableItem{14, "game_start_time", []byte{}},
		},

		// server_query_info is uncompressed and has 1 entry
		{
			"07_50_uncompressed.pbmsg",
			"server_query_info",
			1,
			&stringTableItem{0, "QueryPort", []byte{0x0, 0x0, 0x0, 0x0}},
			&stringTableItem{0, "QueryPort", []byte{0x0, 0x0, 0x0, 0x0}},
		},

		// lightstyles is compressed and has NNNNNNN entries with values
		{
			"05_590_compressed.pbmsg",
			"lightstyles",
			64,
			&stringTableItem{0, "0", []byte{0x6D, 0x00}},
			&stringTableItem{63, "63", []byte{0x61, 0x00}},
		},

		// instancebaseline is compressed and has 75 entries with values
		// XXX TODO: I'm suspicious about the last keys having no values... make sure a delta matches the update!
		{
			"04_22356_compressed.pbmsg",
			"instancebaseline",
			75,
			&stringTableItem{0, "664", _read_fixture("string_tables/instancebaseline/0000_664_414")},
			&stringTableItem{74, "387", []byte{}},
		},

		// EntityNames is compressed and has 123 entries
		{
			"08_4162_compressed.pbmsg",
			"EntityNames",
			350,
			&stringTableItem{0, "kobold_taskmaster_speed_aura", []byte{}},
			&stringTableItem{349, "item_flask", []byte{}},
		},

		// EntityNames is compressed and has 123 entries
		{
			"13_18726_compressed.pbmsg",
			"ModifierNames",
			1274,
			&stringTableItem{0, "modifier_disabled_invulnerable", []byte{}},
			&stringTableItem{1273, "modifier_item_yasha", []byte{}},
		},

		// EconItems is not compressed and fails on values
		// XXX TODO: I'm suspicious about the last keys having no values... make sure a delta matches the update!
		{
			"16_559_uncompressed.pbmsg",
			"EconItems",
			57,
			&stringTableItem{0, "6498667144", []byte{}},
			&stringTableItem{56, "422364528", []byte{}},
		},

		// GenericPrecache is uncompressed with a fixed data (bit) length
		{
			"02_33_uncompressed.pbmsg",
			"genericprecache",
			1,
			&stringTableItem{0, "", []byte{0x00}},
			&stringTableItem{0, "", []byte{0x00}},
		},
	}

	// Iterate through test scenarios
	for _, s := range scenarios {
		// Load the message from the fixture
		m := &dota.CSVCMsg_CreateStringTable{}
		err := proto.Unmarshal(_read_fixture(_sprintf("string_tables/%s", s.fixturePath)), m)
		if err != nil {
			t.Errorf("unable to decode %s: %s", s.fixturePath, err)
			continue
		}

		// Decompress the data if need be
		buf := m.GetStringData()
		if m.GetDataCompressed() {
			buf, err = unlzss(buf)
			if err != nil {
				t.Errorf("unable to decompress %s: %s", s.fixturePath, err)
				continue
			}
		}

		// Make sure we're looking at the right table
		assert.Equal(s.tableName, m.GetName(), s.tableName)

		// Parse the table data
		items := parseStringTable(buf, m.GetNumEntries(), "", m.GetUserDataFixedSize(), m.GetUserDataSize(), m.GetFlags(), false)

		// Make sure we have the correct number of entries
		assert.Equal(s.itemCount, len(items), s.tableName)

		// Verify the first and last entries
		if s.firstItem != nil {
			assert.Equal(s.firstItem, items[0], s.tableName)
		}
		if s.lastItem != nil {
			assert.Equal(s.lastItem, items[len(items)-1], s.tableName)
		}
	}
}

func TestParseStringTableUpdate(t *testing.T) {
	assert := assert.New(t)
	buf := _read_fixture("string_tables/updates/tick_03960_table_7_items_13_size_208")

	items := parseStringTable(buf, 13, "", false, 0, 0, false)

	assert.Equal(int32(261), items[0].Index)
	assert.Equal("broodmother_spawn_spiderlings", items[0].Key)
	assert.Equal(int32(262), items[1].Index)
	assert.Equal("broodmother_spin_web", items[1].Key)
	assert.Equal(int32(263), items[2].Index)
	assert.Equal("broodmother_incapacitating_bite", items[2].Key)
}
