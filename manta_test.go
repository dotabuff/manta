package manta

import (
	"os"
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func init() {
	debugMode = true
}

func TestDumpEverything(t *testing.T) {
	t.Skip()

	assert := assert.New(t)

	buf := mustGetReplayData("1560315800", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem")
	parser, err := NewParser(buf)
	if err != nil {
		t.Fatalf("unable to instantiate parser: %s", err)
	}

	msgNums := make(map[int32]int)

	parser.Callbacks.all = func(t int32, m proto.Message) {
		dirName, ok := packetNames[t]
		if !ok {
			dirName = _sprintf("type_%d", t)
		}
		dir := _sprintf("everything/1560315800/%s", dirName)
		os.MkdirAll("fixtures/"+dir, 0755)

		if _, ok := msgNums[t]; !ok {
			msgNums[t] = 0
		}
		msgNums[t]++

		fileName := _sprintf("%s/tick_%05d_msg_%05d", dir, parser.Tick, msgNums[t])
		_dump_fixture(fileName, []byte(_sdump(m)))
	}

	err = parser.Start()
	assert.Nil(err)
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

		assert.Equal(s.expectCombatLogDamage, gotCombatLogDamage, s.matchId)
		assert.Equal(s.expectCombatLogHealing, gotCombatLogHealing, s.matchId)
		assert.Equal(s.expectCombatLogDeaths, gotCombatLogDeaths, s.matchId)
		assert.Equal(s.expectCombatLogEvents, gotCombatLogEvents, s.matchId)
		assert.Equal(s.expectUnitOrderEvents, gotUnitOrderEvents, s.matchId)
	}
}
