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

	lastChatTick := uint32(0)
	lastChatMessage := ""

	parser.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
		lastChatTick = parser.Tick
		lastChatMessage = m.GetParam2()
		return nil
	})

	err = parser.Start()
	assert.Nil(err)

	assert.Equal(uint32(105819), lastChatTick)
	assert.Equal("yah totally", lastChatMessage)
}
