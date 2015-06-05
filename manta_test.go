package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/stretchr/testify/assert"
)

// Simply tests that we can read the outer structure of a real match
func TestOuterParserRealMatch(t *testing.T) {
	assert := assert.New(t)

	parser := NewParserFromFile("replays/real_match.dem")

	tick := uint32(0)
	lastChatTick := uint32(0)
	lastChatMessage := ""

	parser.Callbacks.OnCNETMsg_Tick = func(m *dota.CNETMsg_Tick) error {
		tick = m.GetTick()
		return nil
	}

	parser.Callbacks.OnCUserMessageSayText2 = func(m *dota.CUserMessageSayText2) error {
		lastChatTick = tick
		lastChatMessage = m.GetParam2()
		return nil
	}

	parser.Start()

	assert.Equal(uint32(105819), lastChatTick)
	assert.Equal("yah totally", lastChatMessage)
}
