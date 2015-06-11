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
		{
			"17_335_uncompressed.pbmsg",
			"CombatLogNames",
			24,
			&stringTableItem{0, "dota_unknown", []byte{}},
			&stringTableItem{23, "item_flask", []byte{}},
		},

		// PASSING: downloadables is uncompressed and has no entries (working)
		{
			"01_29_uncompressed.pbmsg",
			"downloadables",
			0,
			nil,
			nil,
		},

		// PASSING: ResponseKeys is uncompressed and has 15 entries (working)
		{
			"18_175_uncompressed.pbmsg",
			"ResponseKeys",
			15,
			&stringTableItem{0, "concept", []byte{}},
			&stringTableItem{14, "game_start_time", []byte{}},
		},

		// PASSING: server_query_info is uncompressed and has 1 entry
		{
			"07_50_uncompressed.pbmsg",
			"server_query_info",
			1,
			&stringTableItem{0, "QueryPort", []byte{0x0, 0x0, 0x0, 0x0}},
			&stringTableItem{0, "QueryPort", []byte{0x0, 0x0, 0x0, 0x0}},
		},

		// FAILING: lightstyles is compressed and has NNNNNNN entries with values
		/*
			{
				"05_590_compressed.pbmsg",
				"lightstyles",
				1,
				&stringTableItem{0, "", []byte{}},
				&stringTableItem{0, "", []byte{}},
			},
		*/

		// FAILING: instancebaseline is compressed and has NNNNNNN entries with values
		/*
			{
				"04_22356_compressed.pbmsg",
				"instancebaseline",
				1,
				&stringTableItem{0, "", []byte{}},
				&stringTableItem{0, "", []byte{}},
			},
		*/

		// PASSING: EntityNames is compressed and has 123 entries
		{
			"08_4162_compressed.pbmsg",
			"EntityNames",
			350,
			&stringTableItem{0, "kobold_taskmaster_speed_aura", []byte{}},
			&stringTableItem{349, "item_flask", []byte{}},
		},

		// PASSING: EntityNames is compressed and has 123 entries
		{
			"13_18726_compressed.pbmsg",
			"ModifierNames",
			1274,
			&stringTableItem{0, "modifier_disabled_invulnerable", []byte{}},
			&stringTableItem{1273, "modifier_item_yasha", []byte{}},
		},

		// FAILING: EconItems is not compressed and fails on values
		/*
			{
				"16_559_uncompressed.pbmsg",
				"EconItems",
				0,
				nil,
				nil,
			},
		*/
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
		assert.Equal(s.tableName, m.GetName())

		// XXX TODO: remove
		t.Logf("name=%s buflen=%d max_entries=%d fixed=%d size=%d size_bits=%d",
			m.GetName(), len(buf), m.GetMaxEntries(), m.GetUserDataFixedSize(), m.GetUserDataSize(), m.GetUserDataSizeBits())

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

func TestStringTableSpecific(t *testing.T) {
	assert := assert.New(t)
	m := &dota.CSVCMsg_CreateStringTable{}
	proto.Unmarshal(_read_fixture("string_tables/08_4162_compressed.pbmsg"), m)

	buf, err := unlzss(m.GetStringData())
	assert.Nil(err)
	items := parseStringTable(buf, 350, false, 0, 0)
	assert.Equal("kobold_taskmaster_speed_aura", items[0].key)
	assert.Equal("gnoll_assassin_envenomed_weapon", items[1].key)
	assert.Equal("forest_troll_high_priest_heal", items[2].key)
	assert.Equal("forest_troll_high_priest_mana_aura", items[3].key)
	assert.Equal("ghost_frost_attack", items[4].key)
	assert.Equal("harpy_storm_chain_lightning", items[5].key)
	assert.Equal("ogre_magi_frost_armor", items[6].key)
	assert.Equal("giant_wolf_critical_strike", items[7].key)
	assert.Equal("alpha_wolf_critical_strike", items[8].key)
	assert.Equal("alpha_wolf_command_aura", items[9].key)
	assert.Equal("mud_golem_hurl_boulder", items[10].key)
	assert.Equal("mud_golem_rock_destroy", items[11].key)
	assert.Equal("satyr_trickster_purge", items[12].key)
	assert.Equal("satyr_soulstealer_mana_burn", items[13].key)
	assert.Equal("centaur_khan_war_stomp", items[14].key)
}
