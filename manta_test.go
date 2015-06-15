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
func TestOuterParserRealMatch(t *testing.T) {
	assert := assert.New(t)

	buf := mustGetReplayData("real_match", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/real_match.dem")

	parser, err := NewParser(buf)
	if err != nil {
		t.Fatal(err)
	}

	totalCombatLogDamage := 0
	totalCombatLogHealing := 0
	numCombatLogEvents := 0
	lastChatTick := uint32(0)
	lastChatMessage := ""

	parser.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
		lastChatTick = parser.Tick
		lastChatMessage = m.GetParam2()
		return nil
	})

	parser.GameEvents.OnDotaCombatlog(func(m *GameEventDotaCombatlog) error {
		numCombatLogEvents += 1

		switch m.Type {
		case 0:
			totalCombatLogDamage += int(m.Value)
		case 1:
			totalCombatLogHealing += int(m.Value)
		}

		return nil
	})

	err = parser.Start()
	assert.Nil(err)

	assert.Equal(1400664, totalCombatLogDamage)
	assert.Equal(62031, totalCombatLogHealing)
	assert.Equal(58776, numCombatLogEvents)
	assert.Equal(uint32(105819), lastChatTick)
	assert.Equal("yah totally", lastChatMessage)
}
