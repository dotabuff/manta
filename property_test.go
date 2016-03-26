package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

const (
	HANDLE_NONE = uint32(16777215)
)

func TestReadProperties(t *testing.T) {
	assert := assert.New(t)

	scenarios := []struct {
		tableName   string                 // the name of the table, must have a sendtable fixture.
		run         bool                   // whether or not we run the test.
		debug       bool                   // whether or not we print debugging output.
		expectCount int                    // how many result entries we expect.
		expectKeys  map[string]interface{} // a map of entries we expect to be present.
	}{

		/*
			class id=571 name=CDOTATeam
			table id=586 name=CDOTATeam version=0
			 -> prop 0: type:uint8(727) name:m_iTeamNum(784) sendNode: (root)(734)
			 -> prop 1: type:CUtlVector< CHandle< CBasePlayer > >(1592) name:m_aPlayers(1593) sendNode: (root)(734)
			 -> prop 2: type:int32(706) name:m_iScore(1594) sendNode: (root)(734)
			 -> prop 3: type:int32(706) name:m_iRoundsWon(1595) sendNode: (root)(734)
			 -> prop 4: type:char[129](1596) name:m_szTeamname(1597) sendNode: (root)(734)
			 -> prop 5: type:int32(706) name:m_iHeroKills(1598) sendNode: (root)(734)
			 -> prop 6: type:int32(706) name:m_iTowerKills(1286) sendNode: (root)(734)
			 -> prop 7: type:int32(706) name:m_iBarracksKills(1599) sendNode: (root)(734)
			 -> prop 8: type:uint32(748) name:m_unTournamentTeamID(1600) sendNode: (root)(734)
			 -> prop 9: type:uint64(715) name:m_ulTeamLogo(1601) sendNode: (root)(734)
			 -> prop 10: type:uint64(715) name:m_ulTeamBaseLogo(1602) sendNode: (root)(734)
			 -> prop 11: type:uint64(715) name:m_ulTeamBannerLogo(1603) sendNode: (root)(734)
			 -> prop 12: type:bool(723) name:m_bTeamComplete(1604) sendNode: (root)(734)
			 -> prop 13: type:Color(774) name:m_CustomHealthbarColor(1605) sendNode: (root)(734)
			 -> prop 14: type:char[33](1606) name:m_szTag(1607) sendNode: (root)(734)
		*/
		{
			tableName:   "CDOTATeam",
			run:         true,
			debug:       false,
			expectCount: 15,
			expectKeys: map[string]interface{}{
				"m_iTeamNum":             uint64(0),
				"m_aPlayers":             uint32(0), // Most certainly wrong
				"m_iScore":               int32(0),
				"m_iRoundsWon":           int32(0),
				"m_szTeamname":           "Unassigned",
				"m_iHeroKills":           int32(0),
				"m_iTowerKills":          int32(0),
				"m_iBarracksKills":       int32(0),
				"m_unTournamentTeamID":   uint64(0),
				"m_ulTeamLogo":           uint64(0),
				"m_ulTeamBaseLogo":       uint64(0),
				"m_ulTeamBannerLogo":     uint64(0),
				"m_bTeamComplete":        false,
				"m_CustomHealthbarColor": int32(0),
				"m_szTag":                "",
			},
		},

		/*
			class id=559 name=CDOTAFogOfWarWasVisible
			table id=568 name=CDOTAFogOfWarWasVisible version=0
			 -> prop 0: type:uint64[1024](1355) name:m_bWasVisible(1356) sendNode: (root)(734)
		*/
		{
			tableName:   "CDOTAFogOfWarWasVisible",
			run:         true,
			debug:       false,
			expectCount: 1024,
			expectKeys: map[string]interface{}{
				"m_bWasVisible.0000": uint64(0),
				"m_bWasVisible.0001": uint64(0),
				"m_bWasVisible.0511": uint64(0),
				"m_bWasVisible.1023": uint64(0),
				"m_bWasVisible.1024": nil,
			},
		},

		/*
			class id=643 name=CRagdollManager
			table id=661 name=CRagdollManager version=0
			 -> prop 0: type:int8(718) name:m_iCurrentMaxRagdollCount(1982) sendNode: (root)(734)
		*/
		{
			tableName:   "CRagdollManager",
			run:         true,
			debug:       false,
			expectCount: 1,
			expectKeys: map[string]interface{}{
				"m_iCurrentMaxRagdollCount": int32(-1),
			},
		},

		/*
			class id=342 name=CDOTA_DataDire
			table id=351 name=CDOTA_DataDire version=0
			 -> prop 0: type:int32[10](1164) name:m_iReliableGold(1165) sendNode: (root)(734)
			 -> prop 1: type:int32[10](1164) name:m_iUnreliableGold(1166) sendNode: (root)(734)
			 -> prop 2: type:uint8(727) name:m_iTeamNum(784) sendNode: (root)(734)
			 -> prop 3: type:int32[30](1167) name:m_iStartingPositions(1168) sendNode: (root)(734)
			 -> prop 4: type:uint64[256](1169) name:m_bWorldTreeState(1170) sendNode: (root)(734)
		*/
		{
			tableName:   "CDOTA_DataDire",
			run:         true,
			debug:       false,
			expectCount: 428,
			expectKeys: map[string]interface{}{
				"m_vecDataTeam.0000.m_iUnreliableGold": int32(625),
				"m_iTeamNum":                           uint64(3),
				"m_bWorldTreeState.0000":               uint64(18446744073709551615),
				"m_bWorldTreeState.0110":               uint64(18446744073709551615),
				"m_bWorldTreeState.0128":               uint64(0),
				"m_bWorldTreeState.0129":               uint64(0),
				"m_bWorldTreeState.0255":               uint64(0),
			},
		},

		{
			tableName:   "CDOTA_DataRadiant",
			run:         true,
			debug:       false,
			expectCount: 428,
			expectKeys: map[string]interface{}{
				"m_vecDataTeam.0000.m_iUnreliableGold": int32(625),
				"m_iTeamNum":                           uint64(2),
				"m_bWorldTreeState.0000":               uint64(18446744073709551615),
				"m_bWorldTreeState.0110":               uint64(18446744073709551615),
				"m_bWorldTreeState.0128":               uint64(0),
				"m_bWorldTreeState.0129":               uint64(0),
				"m_bWorldTreeState.0255":               uint64(0),
			},
		},

		{
			tableName:   "CDOTA_DataSpectator",
			run:         true,
			debug:       false,
			expectCount: 284,
			expectKeys: map[string]interface{}{
				"m_hPrimaryRune":   HANDLE_NONE,
				"m_hSecondaryRune": HANDLE_NONE,
				"m_iNetWorth.0000": int32(625),
				"m_iNetWorth.0009": int32(625),
			},
		},

		{
			tableName:   "CDOTA_DataCustomTeam",
			run:         true,
			debug:       false,
			expectCount: 258,
			expectKeys: map[string]interface{}{
				"m_iTeamNum":             uint64(6),
				"m_bWorldTreeState.0000": uint64(18446744073709551615),
				"m_bWorldTreeState.0110": uint64(18446744073709551615),
				"m_bWorldTreeState.0128": uint64(0),
				"m_bWorldTreeState.0129": uint64(0),
				"m_bWorldTreeState.0255": uint64(0),
			},
		},

		{
			tableName:   "CFogController",
			run:         true,
			debug:       false,
			expectCount: 21,
			expectKeys: map[string]interface{}{
				"colorPrimary":         uint64(4294963108),
				"colorSecondary":       uint64(4294963108),
				"colorPrimaryLerpTo":   uint64(4294963108),
				"colorSecondaryLerpTo": uint64(4294963108),
				"enable":               true,
				"blend":                false,
				"m_bNoReflectionFog":   true,
			},
		},

		{
			tableName:   "CDOTASpectatorGraphManagerProxy",
			run:         true,
			debug:       false,
			expectCount: 412,
			expectKeys: map[string]interface{}{
				"CDOTASpectatorGraphManager.m_rgPlayerGraphData.0000": HANDLE_NONE,
				"CDOTASpectatorGraphManager.m_rgDireNetWorth.0063":    int32(0),
				"CDOTASpectatorGraphManager.m_nGoldGraphVersion":      int32(3),
			},
		},

		{
			tableName:   "CSpeechBubbleManager",
			run:         true,
			debug:       false,
			expectCount: 25, // WRONG, should carry the array index instead of overwriting. Needs some rework
			expectKeys: map[string]interface{}{
				"m_SpeechBubbles.0000.m_hNPC":       HANDLE_NONE,
				"m_SpeechBubbles.0000.m_flDuration": float32(0),
				"m_SpeechBubbles.0000.m_unOffsetX":  uint32(0),
				"m_SpeechBubbles.0000.m_unCount":    uint32(0),
			},
		},

		{
			tableName:   "CBaseEntity",
			run:         true,
			debug:       false,
			expectCount: 35,
			expectKeys: map[string]interface{}{
				"m_pEntity": true,
			},
		},

		{
			tableName:   "CBaseAnimating",
			run:         true,
			debug:       false,
			expectCount: 111,
			expectKeys: map[string]interface{}{
				"CRenderComponent":                      1,
				"CPhysicsComponent":                     1,
				"m_hEffectEntity":                       HANDLE_NONE,
				"CBodyComponentBaseAnimating.m_hParent": HANDLE_NONE,
				"m_hOwnerEntity":                        HANDLE_NONE,
			},
		},

		{
			tableName:   "CWorld",
			run:         true,
			debug:       false,
			expectCount: 139,
			expectKeys: map[string]interface{}{
				"m_iszOptimizedHeightFieldName": string(""),
				"m_hEffectEntity":               HANDLE_NONE,
				"m_pEntity":                     true,
			},
		},

		{
			tableName:   "CIngameEvent_TI5",
			run:         true,
			debug:       false,
			expectCount: 348,
			expectKeys: map[string]interface{}{
				"m_CompendiumChallengeInfo.0023.nType": 0,
			},
		},

		{
			tableName:   "CDOTAPlayer",
			run:         true,
			debug:       false,
			expectCount: 136,
			expectKeys: map[string]interface{}{
				"m_hViewEntity": HANDLE_NONE,
			},
		},
	}

	// Load our send tables
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables/1731962898.pbmsg"), m); err != nil {
		panic(err)
	}

	// Retrieve the flattened field serializer
	p := &Parser{}
	fs := p.parseSendTables(m, newPropertySerializerTable())

	// Iterate through scenarios
	for _, s := range scenarios {
		// Load up a fixture
		buf := _read_fixture(_sprintf("instancebaseline/1731962898_%s.rawbuf", s.tableName))

		serializer := fs.Serializers[s.tableName][0]
		assert.NotNil(serializer)

		// Optionally skip
		if !s.run {
			continue
		}

		// Read properties
		r := newReader(buf)
		props := NewProperties()
		props.readProperties(r, serializer)
		assert.Equal(s.expectCount, len(props.KV))

		for k, v := range s.expectKeys {
			got, _ := props.Fetch(k)
			assert.EqualValues(v, got)
		}

		// There shouldn't be more than 8 bits left in the buffer
		_debugf("Remaining bits %v", r.remBits())
		assert.True(r.remBits() < 8)
	}
}
