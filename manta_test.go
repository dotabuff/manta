package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/stretchr/testify/assert"
)

func init() {
	debugMode = true
}

// Simply tests that we can read the outer structure of a real match
func TestParseRealMatches(t *testing.T) {
	assert := assert.New(t)

	scenarios := []struct {
		matchId                string
		replayUrl              string
		expectCombatLogDamage  int
		expectCombatLogHealing int
		expectCombatLogDeaths  int
		expectCombatLogEvents  int
		expectUnitOrderEvents  int
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
	}

	for _, s := range scenarios {
		buf := mustGetReplayData(s.matchId, s.replayUrl)
		parser, err := NewParser(buf)
		if err != nil {
			t.Errorf("unable to instantiate parser: %s", err)
			continue
		}

		gotCombatLogDamage := 0
		gotCombatLogHealing := 0
		gotCombatLogDeaths := 0
		gotCombatLogEvents := 0
		gotUnitOrderEvents := 0

		parser.Callbacks.OnCDOTAUserMsg_SpectatorPlayerUnitOrders(func(m *dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error {
			gotUnitOrderEvents += 1
			return nil
		})

		parser.GameEvents.OnDotaCombatlog(func(m *GameEventDotaCombatlog) error {
			gotCombatLogEvents += 1
			switch dota.DOTA_COMBATLOG_TYPES(m.Type) {
			case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DAMAGE:
				gotCombatLogDamage += int(m.Value)
			case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DEATH:
				gotCombatLogDeaths += 1
			case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_HEAL:
				gotCombatLogHealing += int(m.Value)
			}
			return nil
		})

		err = parser.Start()
		assert.Nil(err, s.matchId)

		/*
			Use this to write out instancebaseline fixtures
			t, _ := parser.stringTables.getTableByName("instancebaseline")
			for _, i := range t.items {
				classId, _ := atoi32(i.key)
				className := parser.classInfo[classId]
				_dump_fixture(_sprintf("instancebaseline/%s_%s.rawbuf", className), s.matchId, i.value)
			}
		*/

		/*
				Use this to dump layout of sendtables
			out := _sprintf("")
			classIds := make([]int, 0)
			for classId, _ := range parser.classInfo {
				classIds = append(classIds, int(classId))
			}
			sort.Ints(classIds)
			for _, classId := range classIds {
				className := parser.classInfo[int32(classId)]
				st := parser.sendTables.tables[className]
				out += _sprintf("class id=%d name=%s\n", classId, className)
				out += _sprintf("table id=%d name=%s version=%d\n", st.index, st.name, st.version)
				for i, p := range st.props {
					out += _sprintf(" -> prop %d: %s\n", i, p.Describe())
				}
				out += "\n"
			}
		*/

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
