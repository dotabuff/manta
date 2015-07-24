package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/stretchr/testify/assert"
)

func init() {
	debugMode = true
}

func TestParseCDemoSaveGames(t *testing.T) {
	assert := assert.New(t)

	buf := mustGetReplayData("1560315800", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem")
	parser, err := NewParser(buf)
	assert.NoError(err)

	saves := []map[string]interface{}{}
	parser.Callbacks.OnCDemoSaveGame(func(s *dota.CDemoSaveGame) error {
		save, err := ParseCDemoSaveGame(s)
		saves = append(saves, save)
		return err
	})
	err = parser.Start()
	assert.NoError(err)

	assert.Equal(len(saves), 11)

	assert.EqualValues(saves[0]["matchid"], 1560315800)
	players0 := saves[0]["Players"].(map[string]interface{})
	players02 := players0["2"].(map[string]interface{})
	assert.EqualValues(players02["name"], "kimee")
	assert.EqualValues(players02["hero"], "npc_dota_hero_doom_bringer")

	assert.EqualValues(saves[1]["matchid"], 1560315800)
	players1 := saves[1]["Players"].(map[string]interface{})
	players12 := players1["2"].(map[string]interface{})
	assert.EqualValues(players12["name"], "kimee")
	assert.EqualValues(players12["hero"], "npc_dota_hero_doom_bringer")
}
