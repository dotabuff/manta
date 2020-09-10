package manta

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityOpFlag(t *testing.T) {
	assert := assert.New(t)

	assert.True(EntityOpCreated.Flag(EntityOpCreated))
	assert.False(EntityOpCreated.Flag(EntityOpDeleted))
	assert.False(EntityOpCreated.Flag(EntityOpEntered))
	assert.True(EntityOpCreatedEntered.Flag(EntityOpCreated))
	assert.True(EntityOpCreatedEntered.Flag(EntityOpEntered))
	assert.False(EntityOpCreatedEntered.Flag(EntityOpDeleted))
	assert.False(EntityOpCreatedEntered.Flag(EntityOpLeft))
}

type expectedEntity struct {
	class    string
	property string
	value    interface{}
}

type entityTestScenario struct {
	matchId  string
	tick     uint32
	entities []expectedEntity
	skipInCI bool
}

var entityTestScenarios = map[int64]entityTestScenario{
	3534483793: {
		matchId: "3534483793",
		tick:    22000,
		entities: []expectedEntity{
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerTeamData.0001.m_nSelectedHeroID",
				value:    int32(21),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerTeamData.0001.m_iKills",
				value:    int32(4),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerTeamData.0001.m_hSelectedHero",
				value:    uint64(11961313),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerData.0001.m_bIsValid",
				value:    true,
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerData.0001.m_iPlayerTeam",
				value:    int32(2),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerData.0001.m_iPlayerSteamID",
				value:    uint64(76561198140280423),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerData.0002.m_iszPlayerName",
				value:    "ARMYANO",
			},
		},
	},
	3619005274: {
		matchId: "3619005274",
		tick:    2,
		entities: []expectedEntity{
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerTeamData.0000.m_nSelectedHeroID",
				value:    int32(94),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerTeamData.0000.m_iKills",
				value:    int32(0),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerTeamData.0000.m_hSelectedHero",
				value:    uint64(8750018),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerData.0000.m_bIsValid",
				value:    true,
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerData.0000.m_iPlayerTeam",
				value:    int32(2),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerData.0000.m_iPlayerSteamID",
				value:    uint64(76561198134146057),
			},
			{
				class:    "CDOTA_PlayerResource",
				property: "m_vecPlayerData.0002.m_iszPlayerName",
				value:    "@dogf1ghts", // are very wrong
			},
		},
	},
}

func TestEntities3534483793(t *testing.T) { entityTestScenarios[3534483793].test(t) }

func TestEntities3619005274(t *testing.T) { entityTestScenarios[3619005274].test(t) }

func (s entityTestScenario) test(t *testing.T) {
	assert := assert.New(t)

	if s.skipInCI && os.Getenv("CI") != "" {
		t.Skip("skipping scenario in CI environment")
	}

	r := mustGetReplayReader(s.matchId)
	defer r.Close()

	parser, err := NewStreamParser(r)
	if !assert.Nil(err, "unable to instantiate parser: %s", err) {
		return
	}

	parser.parseToTick(s.tick)

	err = parser.Start()
	if !assert.Nil(err, s.matchId) {
		return
	}

	for _, ee := range s.entities {
		found := false
		for _, ae := range parser.entities {
			if ae == nil || ae.class == nil {
				continue
			}
			if ae.class.name == ee.class {
				found = true
				assert.Equal(ee.value, ae.Get(ee.property),
					"unexpected value for class %s property %s at tick %d", ee.class, ee.property, s.tick)
				break
			}
		}
		assert.True(found, "unable to find entity %s at tick %d", ee.class, s.tick)
	}
}
