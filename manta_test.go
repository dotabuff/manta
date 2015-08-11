package manta

import (
	"sort"
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/stretchr/testify/assert"
)

func init() {
	debugMode = true
}

func TestParseOneMatch(t *testing.T) {
	assert := assert.New(t)

	buf := mustGetReplayData("1560315800", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem")
	parser, err := NewParser(buf)
	assert.Nil(err)
	err = parser.Start()
	assert.Nil(err)
}

// Simply tests that we can read the outer structure of a real match
func TestParseRealMatches(t *testing.T) {
	assert := assert.New(t)

	// Important: make sure to include the most recent test last. These generate fixtures to easily
	// detect diffs in various data structures upon commit, and the latest replay should always be
	// last to provide a most up-to-date baseline.
	scenarios := []struct {
		matchId                string
		replayUrl              string
		expectCombatLogDamage  int32
		expectCombatLogHealing int32
		expectCombatLogDeaths  int32
		expectCombatLogEvents  int32
		expectUnitOrderEvents  int32
	}{
		{
			matchId:                "1560289528",
			replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560289528.dem",
			expectCombatLogDamage:  1180993,
			expectCombatLogHealing: 57511,
			expectCombatLogDeaths:  1449,
			expectCombatLogEvents:  49146,
			expectUnitOrderEvents:  63387,
		},
		{
			matchId:                "1560294294",
			replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560294294.dem",
			expectCombatLogDamage:  768154,
			expectCombatLogHealing: 11565,
			expectCombatLogDeaths:  954,
			expectCombatLogEvents:  24535,
			expectUnitOrderEvents:  30657,
		},
		{
			matchId:                "1560315800",
			replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem",
			expectCombatLogDamage:  1332418,
			expectCombatLogHealing: 57874,
			expectCombatLogDeaths:  1645,
			expectCombatLogEvents:  51288,
			expectUnitOrderEvents:  63992,
		},
		{
			matchId:                "1582611189",
			replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1582611189.dem",
			expectCombatLogDamage:  599388,
			expectCombatLogHealing: 28576,
			expectCombatLogDeaths:  930,
			expectCombatLogEvents:  23800,
			expectUnitOrderEvents:  40237,
		},
		{
			matchId:                "1648457986",
			replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1648457986.dem",
			expectCombatLogDamage:  224773,
			expectCombatLogHealing: 5914,
			expectCombatLogDeaths:  466,
			expectCombatLogEvents:  10170,
			expectUnitOrderEvents:  17822,
		},
	}

	for _, s := range scenarios {
		buf := mustGetReplayData(s.matchId, s.replayUrl)
		parser, err := NewParser(buf)
		if err != nil {
			t.Errorf("unable to instantiate parser: %s", err)
			continue
		}

		gotCombatLogDamage := int32(0)
		gotCombatLogHealing := int32(0)
		gotCombatLogDeaths := int32(0)
		gotCombatLogEvents := int32(0)
		gotUnitOrderEvents := int32(0)

		parser.Callbacks.OnCDOTAUserMsg_SpectatorPlayerUnitOrders(func(m *dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error {
			gotUnitOrderEvents += 1
			return nil
		})

		parser.OnGameEvent("dota_combatlog", func(m *GameEvent) error {
			gotCombatLogEvents += 1

			t, err := m.GetInt32("type")
			assert.Nil(err)

			switch dota.DOTA_COMBATLOG_TYPES(t) {
			case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DAMAGE:
				v, err := m.GetInt32("value")
				assert.Nil(err)
				gotCombatLogDamage += v
			case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DEATH:
				gotCombatLogDeaths += 1
			case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_HEAL:
				v, err := m.GetInt32("value")
				assert.Nil(err)
				gotCombatLogHealing += v
			}

			return nil
		})

		if fixturesMode {
			// Writes out the source_1_legacy_game_events_list.json fixture so that we can identify changes to schema.
			parser.Callbacks.OnCMsgSource1LegacyGameEventList(func(m *dota.CMsgSource1LegacyGameEventList) error {
				_dump_fixture("legacy_game_events/list_latest.json", _json_marshal(m))
				return nil
			})
		}

		err = parser.Start()
		assert.Nil(err, s.matchId)

		if fixturesMode {
			// Use this to write out instancebaseline fixtures
			t, _ := parser.StringTables.GetTableByName("instancebaseline")
			for _, i := range t.Items {
				classId, _ := atoi32(i.Key)
				className := parser.ClassInfo[classId]
				_dump_fixture(_sprintf("instancebaseline/%s_%s.rawbuf", className, s.matchId), i.Value)
			}

			// Use this to dump layout of sendtables
			out := _sprintf("")
			classIds := make([]int, 0)
			for classId, _ := range parser.ClassInfo {
				classIds = append(classIds, int(classId))
			}
			sort.Ints(classIds)
			for _, classId := range classIds {
				className := parser.ClassInfo[int32(classId)]
				st := parser.SendTables.Tables[className]
				out += _sprintf("class id=%d name=%s\n", classId, className)
				out += _sprintf("table id=%d name=%s version=%d\n", st.Index, st.Name, st.Version)
				for i, p := range st.Props {
					out += _sprintf(" -> prop %d: %s\n", i, p.Describe())
				}
				out += "\n"
			}
		}

		assert.Equal(s.expectCombatLogDamage, gotCombatLogDamage, s.matchId)
		assert.Equal(s.expectCombatLogHealing, gotCombatLogHealing, s.matchId)
		assert.Equal(s.expectCombatLogDeaths, gotCombatLogDeaths, s.matchId)
		assert.Equal(s.expectCombatLogEvents, gotCombatLogEvents, s.matchId)
		assert.Equal(s.expectUnitOrderEvents, gotUnitOrderEvents, s.matchId)
	}
}

func BenchmarkParseMatch(b *testing.B) {
	assert := assert.New(b)

	buf := mustGetReplayData("1560315800", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		parser, err := NewParser(buf)
		assert.Nil(err)
		err = parser.Start()
		assert.Nil(err)
	}

	b.ReportAllocs()
}
