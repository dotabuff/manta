package manta

import (
	"os"
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/stretchr/testify/assert"
)

func BenchmarkMatch2159568145(b *testing.B) { testScenarios[2159568145].bench(b) }

// Test client
func TestMatch6682694(t *testing.T) { testScenarios[6682694].test(t) }

func TestMatch5129306977(t *testing.T) { testScenarios[5129306977].test(t) }
func TestMatch5129281647(t *testing.T) { testScenarios[5129281647].test(t) }
func TestMatch4259518439(t *testing.T) { testScenarios[4259518439].test(t) }
func TestMatch4257655794(t *testing.T) { testScenarios[4257655794].test(t) }
func TestMatch3949386909(t *testing.T) { testScenarios[3949386909].test(t) }
func TestMatch3777736409(t *testing.T) { testScenarios[3777736409].test(t) }
func TestMatch3534483793(t *testing.T) { testScenarios[3534483793].test(t) }
func TestMatch3220517753(t *testing.T) { testScenarios[3220517753].test(t) }
func TestMatch2369359192(t *testing.T) { testScenarios[2369359192].test(t) }
func TestMatch2246960647(t *testing.T) { testScenarios[2246960647].test(t) }
func TestMatch2159568145(t *testing.T) { testScenarios[2159568145].test(t) }
func TestMatch2109130988(t *testing.T) { testScenarios[2109130988].test(t) }
func TestMatch1855408730(t *testing.T) { testScenarios[1855408730].test(t) }
func TestMatch1855345768(t *testing.T) { testScenarios[1855345768].test(t) }
func TestMatch1855304265(t *testing.T) { testScenarios[1855304265].test(t) }
func TestMatch1788648401(t *testing.T) { testScenarios[1788648401].test(t) }
func TestMatch1786687320(t *testing.T) { testScenarios[1786687320].test(t) }
func TestMatch1785937100(t *testing.T) { testScenarios[1785937100].test(t) }
func TestMatch1785899023(t *testing.T) { testScenarios[1785899023].test(t) }
func TestMatch1785874713(t *testing.T) { testScenarios[1785874713].test(t) }
func TestMatch1781640270(t *testing.T) { testScenarios[1781640270].test(t) }
func TestMatch1764592109(t *testing.T) { testScenarios[1764592109].test(t) }
func TestMatch1763193771(t *testing.T) { testScenarios[1763193771].test(t) }
func TestMatch1763177231(t *testing.T) { testScenarios[1763177231].test(t) }
func TestMatch1734886116(t *testing.T) { testScenarios[1734886116].test(t) }
func TestMatch1731962898(t *testing.T) { testScenarios[1731962898].test(t) }
func TestMatch1716444111(t *testing.T) { testScenarios[1716444111].test(t) }
func TestMatch1712853372(t *testing.T) { testScenarios[1712853372].test(t) }
func TestMatch1648457986(t *testing.T) { testScenarios[1648457986].test(t) }
func TestMatch1605340040(t *testing.T) { testScenarios[1605340040].test(t) }
func TestMatch1582611189(t *testing.T) { testScenarios[1582611189].test(t) }
func TestMatch1560315800(t *testing.T) { testScenarios[1560315800].test(t) }
func TestMatch1560294294(t *testing.T) { testScenarios[1560294294].test(t) }
func TestMatch1560289528(t *testing.T) { testScenarios[1560289528].test(t) }

type testScenario struct {
	matchId                  string
	replayUrl                string
	debugLevel               uint
	debugTick                uint32
	skipPacketEntities       bool
	expectGameBuild          uint32
	expectEntityEvents       int32
	expectCombatLogDamage    int32
	expectCombatLogHealing   int32
	expectCombatLogDeaths    int32
	expectCombatLogEvents    int32
	expectUnitOrderEvents    int32
	expectPlayer6Name        string
	expectPlayer6Steamid     uint64
	expectHeroEntityName     string
	expectHeroEntityMana     float32
	expectHeroEntityPlayerId int32
	skipInCI                 bool
}

var testScenarios = map[int64]testScenario{
	5129306977: {
		matchId:                  "5129306977",
		replayUrl:                "https://s3-us-west-2.amazonaws.com/manta.dotabuff/5129306977.dem",
		expectGameBuild:          3846,
		expectEntityEvents:       2466090,
		expectUnitOrderEvents:    36185,
		expectHeroEntityName:     "CDOTA_Unit_Hero_Warlock",
		expectHeroEntityMana:     2582.94,
		expectHeroEntityPlayerId: 8,
	},
	5129281647: {
		matchId:                  "5129281647",
		replayUrl:                "https://s3-us-west-2.amazonaws.com/manta.dotabuff/5129281647.dem",
		expectGameBuild:          3846,
		expectEntityEvents:       2919107,
		expectUnitOrderEvents:    43538,
		expectHeroEntityName:     "CDOTA_Unit_Hero_Pudge",
		expectHeroEntityMana:     1647.9391,
		expectHeroEntityPlayerId: 8,
	},
	4259518439: {
		matchId:                  "4259518439",
		replayUrl:                "https://s3-us-west-2.amazonaws.com/manta.dotabuff/4259518439.dem",
		expectGameBuild:          3267,
		expectEntityEvents:       2709121,
		expectUnitOrderEvents:    86120,
		expectHeroEntityName:     "CDOTA_Unit_Hero_Medusa",
		expectHeroEntityMana:     1764.9392,
		expectHeroEntityPlayerId: 1,
	},
	4257655794: {
		matchId:                  "4257655794",
		replayUrl:                "https://s3-us-west-2.amazonaws.com/manta.dotabuff/4257655794.dem",
		expectGameBuild:          3262,
		expectEntityEvents:       2319568,
		expectUnitOrderEvents:    78510,
		expectHeroEntityName:     "CDOTA_Unit_Hero_Dazzle",
		expectHeroEntityMana:     1734.9392,
		expectHeroEntityPlayerId: 6,
	},
	3949386909: {
		matchId:                  "3949386909",
		replayUrl:                "https://s3-us-west-2.amazonaws.com/manta.dotabuff/3949386909.dem",
		expectGameBuild:          2956,
		expectEntityEvents:       3339557,
		expectUnitOrderEvents:    77975,
		expectHeroEntityName:     "CDOTA_Unit_Hero_Bloodseeker",
		expectHeroEntityMana:     1101.9386,
		expectHeroEntityPlayerId: 9,
	},

	3777736409: {
		matchId:                  "3777736409",
		replayUrl:                "https://s3-us-west-2.amazonaws.com/manta.dotabuff/3777736409.dem",
		expectGameBuild:          2735,
		expectEntityEvents:       2263283,
		expectUnitOrderEvents:    58814,
		expectHeroEntityName:     "CDOTA_Unit_Hero_Lion",
		expectHeroEntityMana:     1106.9386,
		expectHeroEntityPlayerId: 6,
	},

	3534483793: {
		matchId:                  "3534483793",
		replayUrl:                "https://s3-us-west-2.amazonaws.com/manta.dotabuff/3534483793.dem",
		expectGameBuild:          2463,
		expectEntityEvents:       2170677,
		expectUnitOrderEvents:    44582,
		expectHeroEntityName:     "CDOTA_Unit_Hero_Rattletrap",
		expectHeroEntityMana:     1293.9387,
		expectHeroEntityPlayerId: 8,
	},

	6682694: {
		matchId:               "6682694",
		replayUrl:             "https://s3-us-west-2.amazonaws.com/manta.dotabuff/6682694.dem",
		expectGameBuild:       1083,
		expectEntityEvents:    3586579,
		expectUnitOrderEvents: 67817,
	},

	3220517753: {
		matchId:                "3220517753",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/3220517753.dem",
		expectGameBuild:        2163,
		expectEntityEvents:     4624363,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		expectUnitOrderEvents:  108879,
	},

	2369359192: {
		matchId:                "2369359192",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/2369359192.dem",
		expectGameBuild:        1449,
		expectEntityEvents:     2434738,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		expectUnitOrderEvents:  64186,
	},

	2246960647: {
		matchId:                "2246960647",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/2246960647.dem",
		expectGameBuild:        1339,
		expectEntityEvents:     4455033,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		expectUnitOrderEvents:  42687,
	},

	2159568145: {
		matchId:                "2159568145",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/2159568145.dem",
		expectGameBuild:        1295,
		expectEntityEvents:     1831423,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		expectUnitOrderEvents:  27202,
	},

	2109130988: {
		matchId:                "2109130988",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/2109130988.dem",
		expectGameBuild:        1253,
		expectEntityEvents:     1072803,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		expectUnitOrderEvents:  40297,
	},

	1855408730: {
		matchId:                "1855408730",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1855408730.dem",
		expectGameBuild:        1106,
		expectEntityEvents:     1397062,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		expectUnitOrderEvents:  34862,
	},
	1855345768: {
		matchId:                "1855345768",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1855345768.dem",
		expectGameBuild:        1104,
		expectEntityEvents:     1567910,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		expectUnitOrderEvents:  41874,
		skipInCI:               true,
	},

	1855304265: {
		matchId:                "1855304265",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1855304265.dem",
		expectGameBuild:        1101,
		expectEntityEvents:     2477532,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		expectUnitOrderEvents:  64823,
		skipInCI:               true,
	},
	1788648401: {
		matchId:                  "1788648401",
		replayUrl:                "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1788648401.dem",
		expectGameBuild:          1036,
		expectEntityEvents:       2357365,
		expectCombatLogDamage:    0,
		expectCombatLogHealing:   0,
		expectCombatLogDeaths:    0,
		expectCombatLogEvents:    0,
		expectHeroEntityName:     "CDOTA_Unit_Hero_Earthshaker",
		expectHeroEntityMana:     1189.9386,
		expectHeroEntityPlayerId: 8,
	},
	1786687320: {
		matchId:                "1786687320",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1786687320.dem",
		expectGameBuild:        1033,
		expectCombatLogDamage:  0,
		expectCombatLogHealing: 0,
		expectCombatLogDeaths:  0,
		expectCombatLogEvents:  0,
		skipInCI:               true,
	},
	1785937100: {
		matchId:                "1785937100",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1785937100.dem",
		expectGameBuild:        1033,
		expectEntityEvents:     1965109,
		expectCombatLogDamage:  955729,
		expectCombatLogHealing: 33158,
		expectCombatLogDeaths:  1345,
		expectCombatLogEvents:  41529,
		expectUnitOrderEvents:  52359,
		expectPlayer6Name:      "JiimoxD",
		expectPlayer6Steamid:   76561198203594628,
		skipInCI:               true,
	},
	1785899023: {
		matchId:                "1785899023",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1785899023.dem",
		expectGameBuild:        1033,
		expectEntityEvents:     2419045,
		expectCombatLogDamage:  1803248,
		expectCombatLogHealing: 48337,
		expectCombatLogDeaths:  1954,
		expectCombatLogEvents:  78804,
		expectUnitOrderEvents:  58269,
		expectPlayer6Name:      "+27",
		expectPlayer6Steamid:   76561198063151170,
		skipInCI:               true,
	},
	1785874713: {
		matchId:                "1785874713",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1785874713.dem",
		expectGameBuild:        1033,
		expectEntityEvents:     1381012,
		expectCombatLogDamage:  513912,
		expectCombatLogHealing: 33359,
		expectCombatLogDeaths:  749,
		expectCombatLogEvents:  21840,
		expectUnitOrderEvents:  40240,
		expectPlayer6Name:      "San-Say",
		expectPlayer6Steamid:   76561198020188611,
		skipInCI:               true,
	},
	1781640270: {
		matchId:                "1781640270",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1781640270.dem",
		expectGameBuild:        1027,
		expectEntityEvents:     1057562,
		expectCombatLogDamage:  345422,
		expectCombatLogHealing: 25884,
		expectCombatLogDeaths:  645,
		expectCombatLogEvents:  18313,
		expectUnitOrderEvents:  39222,
		expectPlayer6Name:      "GGGGGGGGGG",
		expectPlayer6Steamid:   76561198032710514,
	},
	1764592109: {
		matchId:                "1764592109",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1764592109.dem",
		expectGameBuild:        1017,
		expectEntityEvents:     1827933,
		expectCombatLogDamage:  1008784,
		expectCombatLogHealing: 33784,
		expectCombatLogDeaths:  1631,
		expectCombatLogEvents:  42228,
		expectUnitOrderEvents:  36172,
		expectPlayer6Name:      "Doffo",
		expectPlayer6Steamid:   76561198087353732,
		skipInCI:               true,
	},
	1763193771: {
		matchId:                "1763193771",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1763193771.dem",
		expectGameBuild:        1016,
		expectEntityEvents:     1203640,
		expectCombatLogDamage:  623594,
		expectCombatLogHealing: 19530,
		expectCombatLogDeaths:  1022,
		expectCombatLogEvents:  24436,
		expectUnitOrderEvents:  31994,
		expectPlayer6Name:      "Monst_er",
		expectPlayer6Steamid:   76561198201328510,
		skipInCI:               true,
	},
	1763177231: {
		matchId:                "1763177231",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1763177231.dem",
		expectGameBuild:        1016,
		expectEntityEvents:     1221874,
		expectCombatLogDamage:  479350,
		expectCombatLogHealing: 22611,
		expectCombatLogDeaths:  977,
		expectCombatLogEvents:  20043,
		expectUnitOrderEvents:  35975,
		expectPlayer6Name:      "Supercowman",
		expectPlayer6Steamid:   76561198013311415,
		skipInCI:               true,
	},
	1734886116: {
		matchId:                "1734886116",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1734886116.dem",
		expectGameBuild:        1003,
		expectEntityEvents:     2049211,
		expectCombatLogDamage:  1048805,
		expectCombatLogHealing: 25089,
		expectCombatLogDeaths:  1447,
		expectCombatLogEvents:  42307,
		expectUnitOrderEvents:  59775,
		expectPlayer6Name:      "Eggard",
		expectPlayer6Steamid:   76561197972599979,
	},
	1731962898: {
		matchId:                "1731962898",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1731962898.dem",
		expectGameBuild:        1003,
		expectEntityEvents:     1183267,
		expectCombatLogDamage:  415560,
		expectCombatLogHealing: 20018,
		expectCombatLogDeaths:  690,
		expectCombatLogEvents:  24296,
		expectUnitOrderEvents:  27525,
		expectPlayer6Name:      "Snayp8",
		expectPlayer6Steamid:   76561198047587062,
		skipInCI:               true,
	},
	1716444111: {
		matchId:                "1716444111",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1716444111.dem",
		expectGameBuild:        995,
		expectEntityEvents:     2854511,
		expectCombatLogDamage:  1398735,
		expectCombatLogHealing: 49659,
		expectCombatLogDeaths:  2169,
		expectCombatLogEvents:  76921,
		expectUnitOrderEvents:  48822,
		expectPlayer6Name:      "GangBang",
		expectPlayer6Steamid:   76561198136700681,
		skipInCI:               true,
	},
	1712853372: {
		matchId:                "1712853372",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1712853372.dem",
		expectGameBuild:        991,
		expectEntityEvents:     1708696,
		expectCombatLogDamage:  671297,
		expectCombatLogHealing: 23467,
		expectCombatLogDeaths:  1099,
		expectCombatLogEvents:  30381,
		expectUnitOrderEvents:  48107,
		expectPlayer6Name:      "BFG",
		expectPlayer6Steamid:   76561198047707927,
		skipInCI:               true,
	},
	1648457986: {
		matchId:                "1648457986",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1648457986.dem",
		expectGameBuild:        962,
		expectEntityEvents:     742252,
		expectCombatLogDamage:  224773,
		expectCombatLogHealing: 5914,
		expectCombatLogDeaths:  466,
		expectCombatLogEvents:  10170,
		expectUnitOrderEvents:  17822,
		expectPlayer6Name:      "Grinder",
		expectPlayer6Steamid:   76561198207988337,
		skipInCI:               true,
	},
	1605340040: {
		matchId:                "1605340040",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1605340040.dem",
		expectGameBuild:        955,
		expectEntityEvents:     1283574,
		expectCombatLogDamage:  522367,
		expectCombatLogHealing: 31721,
		expectCombatLogDeaths:  795,
		expectCombatLogEvents:  21116,
		expectUnitOrderEvents:  40669,
		expectPlayer6Name:      "Az ★ | Big A Man",
		expectPlayer6Steamid:   76561198156504817,
		expectHeroEntityName:   "CDOTA_Unit_Hero_Chen",
		expectHeroEntityMana:   1159.9386,
		skipInCI:               true,
	},
	1582611189: {
		matchId:                "1582611189",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1582611189.dem",
		expectGameBuild:        944,
		expectEntityEvents:     1427025,
		expectCombatLogDamage:  599388,
		expectCombatLogHealing: 28576,
		expectCombatLogDeaths:  930,
		expectCombatLogEvents:  23800,
		expectUnitOrderEvents:  40237,
		expectPlayer6Name:      "The13ananaMan",
		expectPlayer6Steamid:   76561198068424443,
		skipInCI:               true,
	},
	1560315800: {
		matchId:                "1560315800",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem",
		expectGameBuild:        928,
		expectEntityEvents:     2781076,
		expectCombatLogDamage:  1332418,
		expectCombatLogHealing: 57874,
		expectCombatLogDeaths:  1645,
		expectCombatLogEvents:  51288,
		expectUnitOrderEvents:  63992,
		expectPlayer6Name:      "ariethebeast",
		expectPlayer6Steamid:   76561198065323776,
		expectHeroEntityName:   "CDOTA_Unit_Hero_Pudge",
		expectHeroEntityMana:   858.10474,
		skipInCI:               true,
	},
	1560294294: {
		matchId:                "1560294294",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560294294.dem",
		expectGameBuild:        928,
		expectEntityEvents:     1611898,
		expectCombatLogDamage:  768154,
		expectCombatLogHealing: 11565,
		expectCombatLogDeaths:  954,
		expectCombatLogEvents:  24535,
		expectUnitOrderEvents:  30657,
		expectPlayer6Name:      "Laslo",
		expectPlayer6Steamid:   76561198034549887,
		skipInCI:               true,
	},
	1560289528: {
		matchId:                "1560289528",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560289528.dem",
		expectGameBuild:        928,
		expectEntityEvents:     2270022,
		expectCombatLogDamage:  1180993,
		expectCombatLogHealing: 57511,
		expectCombatLogDeaths:  1449,
		expectCombatLogEvents:  49146,
		expectUnitOrderEvents:  63387,
		expectPlayer6Name:      "It takes a tree to tango",
		expectPlayer6Steamid:   76561197993050562,
		expectHeroEntityName:   "CDOTA_Unit_Hero_Undying",
		expectHeroEntityMana:   1108.1353,
		skipInCI:               true,
	},
}

func (s testScenario) bench(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := mustGetReplayReader(s.matchId, s.replayUrl)

		parser, err := NewStreamParser(r)
		if err != nil {
			b.Fatalf("unable to instantiate parser: %s", err)
		}

		parser.Callbacks.OnCDOTAUserMsg_SpectatorPlayerUnitOrders(func(m *dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error { return nil })
		parser.Callbacks.OnCDemoFileInfo(func(m *dota.CDemoFileInfo) error { return nil })
		parser.OnEntity(func(e *Entity, op EntityOp) error { return nil })
		parser.OnGameEvent("dota_combatlog", func(m *GameEvent) error { return nil })

		if err := parser.Start(); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportAllocs()
}

func (s testScenario) test(t *testing.T) {
	assert := assert.New(t)

	if s.debugLevel > 0 {
		if s.debugTick == 0 {
			debugLevel = s.debugLevel
		}

		defer func() {
			debugLevel = 0
		}()
	}

	if s.skipInCI && os.Getenv("CI") != "" {
		t.Skip("skipping scenario in CI environment")
	}

	t.Parallel()

	r := mustGetReplayReader(s.matchId, s.replayUrl)
	defer r.Close()

	parser, err := NewStreamParser(r)
	if err != nil {
		t.Errorf("unable to instantiate parser: %s", err)
		return
	}

	gotFileInfo := false
	gotCombatLogDamage := int32(0)
	gotCombatLogHealing := int32(0)
	gotCombatLogDeaths := int32(0)
	gotCombatLogEvents := int32(0)
	gotUnitOrderEvents := int32(0)
	gotEntityEvents := int32(0)
	gotPlayer6Name := "<Missing>"
	gotPlayer6Steamid := uint64(0)
	gotHeroEntityMana := float32(0.0)
	gotHeroEntityPlayerId := int32(0)

	if s.debugTick > 0 {
		parser.Callbacks.OnCNETMsg_Tick(func(m *dota.CNETMsg_Tick) error {
			if parser.Tick >= s.debugTick {
				debugLevel = s.debugLevel
			}
			return nil
		})
	}

	parser.Callbacks.OnCDOTAUserMsg_SpectatorPlayerUnitOrders(func(m *dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error {
		gotUnitOrderEvents += 1
		return nil
	})

	parser.Callbacks.OnCDemoFileInfo(func(m *dota.CDemoFileInfo) error {
		gotFileInfo = true
		return nil
	})

	parser.OnEntity(func(e *Entity, op EntityOp) error {
		gotEntityEvents += 1

		if e.class.name == s.expectHeroEntityName {
			if v, ok := e.Get("m_flMaxMana").(float32); ok {
				gotHeroEntityMana = v
			}
			if v, ok := e.GetInt32("m_iPlayerID"); ok {
				gotHeroEntityPlayerId = v
			}
		}

		if e.class.name == "CDOTA_PlayerResource" {
			if v, ok := e.Get("m_vecPlayerData.0006.m_iszPlayerName").(string); ok {
				gotPlayer6Name = v
			} else if v, ok := e.Get("m_iszPlayerNames.0006").(string); ok {
				gotPlayer6Name = v
			}

			if v, ok := e.Get("m_vecPlayerData.0006.m_iPlayerSteamID").(uint64); ok {
				gotPlayer6Steamid = v
			} else if v, ok := e.Get("m_iPlayerSteamIDs.0006").(uint64); ok {
				gotPlayer6Steamid = v
			}
		}

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

	err = parser.Start()
	assert.Nil(err, s.matchId)

	assert.True(gotFileInfo)
	assert.Equal(s.expectGameBuild, parser.GameBuild, s.matchId)
	if s.expectEntityEvents > 0 {
		assert.Equal(s.expectEntityEvents, gotEntityEvents, s.matchId)
	}
	if s.expectCombatLogDamage > 0 {
		assert.Equal(s.expectCombatLogDamage, gotCombatLogDamage, s.matchId)
	}
	if s.expectCombatLogHealing > 0 {
		assert.Equal(s.expectCombatLogHealing, gotCombatLogHealing, s.matchId)
	}
	if s.expectCombatLogDeaths > 0 {
		assert.Equal(s.expectCombatLogDeaths, gotCombatLogDeaths, s.matchId)
	}
	if s.expectCombatLogEvents > 0 {
		assert.Equal(s.expectCombatLogEvents, gotCombatLogEvents, s.matchId)
	}
	if s.expectUnitOrderEvents > 0 {
		assert.Equal(s.expectUnitOrderEvents, gotUnitOrderEvents, s.matchId)
	}
	if s.expectPlayer6Name != "" {
		assert.Equal(s.expectPlayer6Name, gotPlayer6Name, s.matchId)
	}
	if s.expectPlayer6Steamid > 0 {
		assert.Equal(s.expectPlayer6Steamid, gotPlayer6Steamid, s.matchId)
	}
	if s.expectHeroEntityMana > 0.0 {
		assert.Equal(s.expectHeroEntityMana, gotHeroEntityMana, s.matchId)
	}
	if s.expectHeroEntityPlayerId > 0.0 {
		assert.Equal(s.expectHeroEntityPlayerId, gotHeroEntityPlayerId, s.matchId)
	}
}
