package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/stretchr/testify/assert"
)

func TestParseCDemoSaveGames(t *testing.T) {
	assert := assert.New(t)

	buf := mustGetReplayData("1560315800", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem")
	parser, err := NewParser(buf)
	assert.NoError(err)

	saves := []*SaveGame{}
	parser.Callbacks.OnCDemoSaveGame(func(s *dota.CDemoSaveGame) error {
		save, err := ParseCDemoSaveGame(s)
		saves = append(saves, save)
		return err
	})
	err = parser.Start()
	assert.NoError(err)

	assert.Equal(len(saves), 11)

	unitLengths := []int{12, 15, 14, 15, 13, 15, 15, 14, 16, 14, 16}

	for n, save := range saves {
		assert.Len(save.Players, 10, "Players")
		assert.Len(save.Units, unitLengths[n], "Units")
		assert.Len(save.StockInfo, 10, "StockInfo")
		assert.EqualValues(save.Gameid, 0, "Gameid")
		assert.EqualValues(save.Matchid, 1560315800, "Matchid")
		assert.EqualValues(save.Version, 1, "Version")

		player2 := save.Players["2"]
		assert.EqualValues(player2.Name, "kimee")
		assert.EqualValues(player2.Hero, "npc_dota_hero_doom_bringer")
		assert.EqualValues(player2.Sharemask, 0, "Sharemask")
		assert.EqualValues(player2.Steamid, 76561197961237397, "Steamid")
		assert.EqualValues(player2.Team, 2, "Team")
	}
}
