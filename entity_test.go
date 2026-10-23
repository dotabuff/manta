package manta

import (
	"math"
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
	matchId   string
	replayUrl string
	tick      uint32
	entities  []expectedEntity
	skipInCI  bool
}

var entityTestScenarios = map[int64]entityTestScenario{
	3534483793: {
		matchId:   "3534483793",
		replayUrl: "https://s3-us-west-2.amazonaws.com/manta.dotabuff/3534483793.dem",
		tick:      22000,
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
		matchId:   "3619005274",
		replayUrl: "https://s3-us-west-2.amazonaws.com/manta.dotabuff/3619005274.dem",
		tick:      2,
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

	r := mustGetReplayReader(s.matchId, s.replayUrl)
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

// TestEntityTypedGettersMatchGet verifies the typed getters, which read state
// cells directly, agree with their definition in terms of Get (a type
// assertion on the boxed value) for every field of every entity at the end of
// a replay, and for names that do not exist.
func TestEntityTypedGettersMatchGet(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := mustGetReplayReader("8552595443", testScenarios[8552595443].replayUrl)
	defer r.Close()
	p, err := NewStreamParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Start(); err != nil {
		t.Fatal(err)
	}

	checked := 0
	for _, e := range p.FilterEntity(func(*Entity) bool { return true }) {
		names := []string{"m_no_such_field", "m_vecPlayerData.9999.m_iszPlayerName"}
		for name := range e.Map() {
			names = append(names, name)
		}
		for _, name := range names {
			v := e.Get(name)

			wantI32, okI32 := v.(int32)
			gotI32, gotOkI32 := e.GetInt32(name)
			assert.Equal(okI32, gotOkI32, name)
			assert.Equal(wantI32, gotI32, name)

			var wantU32 uint32
			okU32 := true
			switch x := v.(type) {
			case uint32:
				wantU32 = x
			case uint64:
				wantU32 = uint32(x)
			default:
				okU32 = false
			}
			gotU32, gotOkU32 := e.GetUint32(name)
			assert.Equal(okU32, gotOkU32, name)
			assert.Equal(wantU32, gotU32, name)

			wantU64, okU64 := v.(uint64)
			gotU64, gotOkU64 := e.GetUint64(name)
			assert.Equal(okU64, gotOkU64, name)
			assert.Equal(wantU64, gotU64, name)

			wantF32, okF32 := v.(float32)
			gotF32, gotOkF32 := e.GetFloat32(name)
			assert.Equal(okF32, gotOkF32, name)
			assert.Equal(math.Float32bits(wantF32), math.Float32bits(gotF32), name)

			wantStr, okStr := v.(string)
			gotStr, gotOkStr := e.GetString(name)
			assert.Equal(okStr, gotOkStr, name)
			assert.Equal(wantStr, gotStr, name)

			wantBool, okBool := v.(bool)
			gotBool, gotOkBool := e.GetBool(name)
			assert.Equal(okBool, gotOkBool, name)
			assert.Equal(wantBool, gotBool, name)

			assert.Equal(v != nil, e.Exists(name), name)
			checked++
		}
	}
	if checked < 10000 {
		t.Fatalf("only checked %d fields", checked)
	}
}

// requireComparable compiles only for comparable types.
func requireComparable[T comparable]() {}

// TestEntityIsComparable guards that Entity stays a comparable type: callers may
// compare Entity values with ==, use them as map keys, or pass them as a
// comparable type argument, so a slice-holding field would be a breaking API
// change. This fails to compile if Entity stops being comparable.
func TestEntityIsComparable(t *testing.T) {
	requireComparable[Entity]()
}
