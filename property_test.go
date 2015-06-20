package manta

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
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
				// 15 index bits, leaves 1 byte unaccounted for.
				// Based on m_aPlayers being a single bit...
				// manta.(*reader).dumpBits: @ bit 00024 (byte 003 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 16638200126367596544           | bitfloat32: -1.1368684e-13 | string: -
				// manta.(*reader).dumpBits: @ bit 00025 (byte 003 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 8319100063183798272            | bitfloat32: 8.796093e+12 | string: -
				// manta.(*reader).dumpBits: @ bit 00026 (byte 003 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 13382922068446674944           | bitfloat32: 2.2737368e-13 | string: -
				// manta.(*reader).dumpBits: @ bit 00027 (byte 003 + 3)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 6691461034223337472            | bitfloat32: -3.877409e-26 | string: -
				// manta.(*reader).dumpBits: @ bit 00028 (byte 003 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 3345730517111668736            | bitfloat32: -5.24288e+06 | string: -
				// manta.(*reader).dumpBits: @ bit 00029 (byte 003 + 5)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 10896237295410610176           | bitfloat32: -6.1390764e+22 | string: -
				// manta.(*reader).dumpBits: @ bit 00030 (byte 003 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 5448118647705305088            | bitfloat32: 6.6551657e+30 | string: -
				// manta.(*reader).dumpBits: @ bit 00031 (byte 003 + 7)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 11947431360707428352           | bitfloat32: -0.00020217896 | string: -
				"m_iTeamNum": uint8(0),
				// Maybe a single bit to say no data / zero entries?
				// manta.(*reader).dumpBits: @ bit 00032 (byte 004 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 15197087717208489984           | bitfloat32: -3.8280597e+17 | string: -
				"m_aPlayers": uint(0), // Most certainly wrong
				// manta.(*reader).dumpBits: @ bit 00033 (byte 004 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 7598543858604244992            | bitfloat32: 1.6480077e+28 | string: -
				"m_iScore": int32(0),
				// manta.(*reader).dumpBits: @ bit 00041 (byte 005 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 7451613997854250240            | bitfloat32: 2.7477812e+20 | string: -
				"m_iRoundsWon": int32(0),
				// manta.(*reader).dumpBits: @ bit 00049 (byte 006 + 1)  | binary: 1 | uint8: 85  | var32: -43         | varu32: 85         | varu64: 85                   | uint64: 7955443211351191125            | bitfloat32: 1.7860483e+31 | string: Unassigned
				"m_szTeamname": "Unassigned",
				// manta.(*reader).dumpBits: @ bit 00137 (byte 017 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 0                              | bitfloat32: 0            | string: -
				"m_iHeroKills": int32(0),
				// manta.(*reader).dumpBits: @ bit 00145 (byte 018 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 0                              | bitfloat32: 0            | string: -
				"m_iTowerKills": int32(0),
				// manta.(*reader).dumpBits: @ bit 00153 (byte 019 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: 0            | string: -
				"m_iBarracksKills": int32(0),
				// manta.(*reader).dumpBits: @ bit 00161 (byte 020 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: 0            | string: -
				"m_unTournamentTeamID": uint32(0),
				// manta.(*reader).dumpBits: @ bit 00169 (byte 021 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: 0            | string: -
				"m_ulTeamLogo": uint64(0),
				// manta.(*reader).dumpBits: @ bit 00177 (byte 022 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: 0            | string: -
				"m_ulTeamBaseLogo": uint64(0),
				// manta.(*reader).dumpBits: @ bit 00185 (byte 023 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				"m_ulTeamBannerLogo": uint64(0),
				// manta.(*reader).dumpBits: @ bit 00193 (byte 024 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				"m_bTeamComplete": false,
				// The last 2 fields have unknown alignment.
				// manta.(*reader).dumpBits: @ bit 00194 (byte 024 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00195 (byte 024 + 3)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00196 (byte 024 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00197 (byte 024 + 5)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00198 (byte 024 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00199 (byte 024 + 7)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00200 (byte 025 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00201 (byte 025 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00202 (byte 025 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				// manta.(*reader).dumpBits: @ bit 00203 (byte 025 + 3)  | binary: 0 | uint8: 128 | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00204 (byte 025 + 4)  | binary: 0 | uint8: 192 | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00205 (byte 025 + 5)  | binary: 0 | uint8: 224 | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00206 (byte 025 + 6)  | binary: 0 | uint8: 240 | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00207 (byte 025 + 7)  | binary: 0 | uint8: 248 | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00208 (byte 026 + 0)  | binary: 0 | uint8: 252 | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00209 (byte 026 + 1)  | binary: 0 | uint8: ERR | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00210 (byte 026 + 2)  | binary: 1 | uint8: ERR | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00211 (byte 026 + 3)  | binary: 1 | uint8: ERR | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00212 (byte 026 + 4)  | binary: 1 | uint8: ERR | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00213 (byte 026 + 5)  | binary: 1 | uint8: ERR | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00214 (byte 026 + 6)  | binary: 1 | uint8: ERR | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
				// manta.(*reader).dumpBits: @ bit 00215 (byte 026 + 7)  | binary: 1 | uint8: ERR | var32: ERR         | varu32: ERR        | varu64: ERR                  | uint64: ERR                            | bitfloat32: ERR          | string: ERR
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
				"m_bWasVisible.0":    uint64(0),
				"m_bWasVisible.1":    uint64(0),
				"m_bWasVisible.511":  uint64(0),
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
				"m_iCurrentMaxRagdollCount": int8(-1),
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
			This structure actually has 10+10+1+30+256=307 entries.
			The packed value data starts at bit 384 (#385), leaving 307 marker bits and the first 78 bits in the wind.
		*/
		{
			tableName:   "CDOTA_DataDire",
			run:         true,
			debug:       false,
			expectCount: (10 + 10 + 1 + 30 + 256),
			expectKeys: map[string]interface{}{
				// manta.(*reader).dumpBits: @ bit 00384 (byte 048 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | bitfloat32: 0            | string: -
				"m_iReliableGold.0": int32(0),
				"m_iReliableGold.1": int32(0),
				"m_iReliableGold.2": int32(0),
				"m_iReliableGold.3": int32(0),
				"m_iReliableGold.4": int32(0),
				"m_iReliableGold.5": int32(0),
				"m_iReliableGold.6": int32(0),
				"m_iReliableGold.7": int32(0),
				"m_iReliableGold.8": int32(0),
				"m_iReliableGold.9": int32(0),
				// manta.(*reader).dumpBits: @ bit 00464 (byte 058 + 0)  | binary: 0 | uint8: 226 | var32: 625         | varu32: 1250       | varu64: 1250                 | bitfloat32: 5.4416815e-33 | string: -
				"m_iUnreliableGold.0": int32(625),
				"m_iUnreliableGold.1": int32(625),
				"m_iUnreliableGold.2": int32(625),
				"m_iUnreliableGold.3": int32(625),
				"m_iUnreliableGold.4": int32(625),
				"m_iUnreliableGold.5": int32(625),
				"m_iUnreliableGold.6": int32(625),
				"m_iUnreliableGold.7": int32(625),
				"m_iUnreliableGold.8": int32(625),
				"m_iUnreliableGold.9": int32(625),
				//manta.(*reader).dumpBits: @ bit 00624 (byte 078 + 0)  | binary: 1 | uint8: 3   | var32: -2          | varu32: 3          | varu64: 3                    | bitfloat32: 2.3694284e-38 | string: -
				"m_iTeamNum": uint8(3),
				// manta.(*reader).dumpBits: @ bit 00632 (byte 079 + 0)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | uint64: 72340172838076673              | bitfloat32: 2.3694278e-38 | string: -
				"m_iStartingPositions.0": int32(-1),
				// manta.(*reader).dumpBits: @ bit 00864 (byte 108 + 0)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | uint64: 18446744073709551361           | bitfloat32: NaN          | string: -
				"m_iStartingPositions.29": int32(-1),
				// Each tree state full of (mostly) 1's takes 10 bytes.
				// manta.(*reader).dumpBits: @ bit 00872 (byte 109 + 0)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | uint64: 18446744073709551615           | bitfloat32: NaN          | string: -
				"m_bWorldTreeState.0": uint64(18446744073709551615),
				// manta.(*reader).dumpBits: @ bit 00952 (byte 119 + 0)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | uint64: 18446744073709551615           | bitfloat32: NaN          | string: -
				"m_bWorldTreeState.1": uint64(18446744073709551615),
				// manta.(*reader).dumpBits: @ bit 11032 (byte 1379 + 0)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | uint64: 18446744073709551615           | bitfloat32: NaN          | string: -
				"m_bWorldTreeState.126": uint64(18446744073709551615),
				"m_bWorldTreeState.127": uint64(18446744073709551615),
				// That was 1270 bytes (127 entries) worth of 10-byte varint 18446744073709551615's.
				// Now we see 127 more bytes (127 entries) worth of 1-byte varint 0's.
				// manta.(*reader).dumpBits: @ bit 11112 (byte 1389 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 0                              | bitfloat32: 0            | string: -
				"m_bWorldTreeState.128": uint64(0),
				"m_bWorldTreeState.129": uint64(0),
				// manta.(*reader).dumpBits: @ bit 12128 (byte 1516 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				"m_bWorldTreeState.255": uint64(0),
			},
		},

		/*
			WIP:

			{
				tableName:   "CBaseAnimating",
				run:         false,
				debug:       false,
				expectCount: 0,
				expectKeys:  map[string]interface{}{},
			},
			{
				tableName:   "CDOTA_PlayerResource",
				expectCount: 0,
				expectKeys: map[string]interface{}{
					"m_iszPlayerNames.1":  "Thias",
					"m_iszPlayerNames.2":  "Ash",
					"m_iPlayerSteamIDs.1": uint64(76561197996062617),
					"m_iPlayerSteamIDs.2": uint64(76561198046993147),
				},
			},
		*/
	}

	// Load our send tables
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables/1560315800.pbmsg"), m); err != nil {
		panic(err)
	}

	st, err := parseSendTables(m)
	assert.Nil(err)

	// Iterate through scenarios
	for _, s := range scenarios {
		// Load up a fixture
		buf := _read_fixture(_sprintf("instancebaseline/1560315800_%s.rawbuf", s.tableName))
		st, ok := st.getTableByName(s.tableName)
		assert.True(ok)

		// Optionally skip
		if !s.run {
			continue
		}

		// Optionally disable debugging
		if !s.debug {
			debugMode = false
		}

		// Read properties
		r := newReader(buf)
		props := readProperties(r, st)
		assert.Equal(s.expectCount, len(props))
		for k, v := range s.expectKeys {
			assert.Equal(v, props[k], s.tableName+"."+k)
		}

		// Re-enable debugging
		debugMode = true
	}
}

// Run this to see the beginning of each instancebaseline for header analysis.
func TestAnalyzeInstancebaselines(t *testing.T) {
	t.Skip()

	assert := assert.New(t)

	// Load our send tables
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables/1560315800.pbmsg"), m); err != nil {
		panic(err)
	}

	st, err := parseSendTables(m)
	assert.Nil(err)

	onlyFixture := os.Getenv("ONLY_FIXTURE")

	// Iterate through fixtures
	files, _ := filepath.Glob("fixtures/instancebaseline/*.rawbuf")
	for _, f := range files {
		fileName := path.Base(f)
		tableName := strings.Split(strings.SplitN(fileName, "_", 2)[1], ".")[0]
		sendTable, ok := st.getTableByName(tableName)
		assert.True(ok)

		if onlyFixture != "" && onlyFixture != tableName {
			continue
		}

		buf := _read_fixture("instancebaseline/" + fileName)
		r := newReader(buf)

		first1 := -1
		for i := 0; i < r.size; i++ {
			if r.readBoolean() {
				first1 = r.pos - 1
				break
			}
		}
		r.pos = 0
		nDump := first1 + 1 + (8 * 4)
		if os.Getenv("ALL_BITS") != "" {
			nDump = r.size
		}
		_debugf("fixture %s (%d props) has first 1 at %s", colorBold(tableName), len(sendTable.props), colorValue(first1))
		for i := 0; i < len(sendTable.props); i++ {
			if i > 3 {
				break
			}
			_debugf("prop %d: %s", i, sendTable.props[i].Describe())
		}
		r.dumpBits(nDump)

		r.pos = 0
		readProperties(r, sendTable)
	}
}
