package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

const (
	HANDLE_NONE = uint32(2097151)
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
				"m_iTeamNum": uint64(0),
				// Maybe a single bit to say no data / zero entries?
				// manta.(*reader).dumpBits: @ bit 00032 (byte 004 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 15197087717208489984           | bitfloat32: -3.8280597e+17 | string: -
				"m_aPlayers": uint32(0), // Most certainly wrong
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
				"m_unTournamentTeamID": uint64(0),
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
			This structure actually has 10+10+1+30+256=307 entries.
			The packed value data starts at bit 384 (#385), leaving 307 marker bits and the first 78 bits in the wind.
		*/
		{
			tableName:   "CDOTA_DataDire",
			run:         false,
			debug:       false,
			expectCount: (10 + 10 + 1 + 30 + 256),
			expectKeys: map[string]interface{}{
				// manta.(*reader).dumpBits: @ bit 00384 (byte 048 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | bitfloat32: 0            | string: -
				"m_iReliableGold.0000": int32(0),
				"m_iReliableGold.0001": int32(0),
				"m_iReliableGold.0002": int32(0),
				"m_iReliableGold.0003": int32(0),
				"m_iReliableGold.0004": int32(0),
				"m_iReliableGold.0005": int32(0),
				"m_iReliableGold.0006": int32(0),
				"m_iReliableGold.0007": int32(0),
				"m_iReliableGold.0008": int32(0),
				"m_iReliableGold.0009": int32(0),
				// manta.(*reader).dumpBits: @ bit 00464 (byte 058 + 0)  | binary: 0 | uint8: 226 | var32: 625         | varu32: 1250       | varu64: 1250                 | bitfloat32: 5.4416815e-33 | string: -
				"m_iUnreliableGold.0000": int32(625),
				"m_iUnreliableGold.0001": int32(625),
				"m_iUnreliableGold.0002": int32(625),
				"m_iUnreliableGold.0003": int32(625),
				"m_iUnreliableGold.0004": int32(625),
				"m_iUnreliableGold.0005": int32(625),
				"m_iUnreliableGold.0006": int32(625),
				"m_iUnreliableGold.0007": int32(625),
				"m_iUnreliableGold.0008": int32(625),
				"m_iUnreliableGold.0009": int32(625),
				//manta.(*reader).dumpBits: @ bit 00624 (byte 078 + 0)  | binary: 1 | uint8: 3   | var32: -2          | varu32: 3          | varu64: 3                    | bitfloat32: 2.3694284e-38 | string: -
				"m_iTeamNum": uint64(3),
				// manta.(*reader).dumpBits: @ bit 00632 (byte 079 + 0)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | uint64: 72340172838076673              | bitfloat32: 2.3694278e-38 | string: -
				"m_iStartingPositions.0000": int32(-1),
				// manta.(*reader).dumpBits: @ bit 00864 (byte 108 + 0)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | uint64: 18446744073709551361           | bitfloat32: NaN          | string: -
				"m_iStartingPositions.0029": int32(-1),
				// Each tree state full of (mostly) 1's takes 10 bytes.
				// manta.(*reader).dumpBits: @ bit 00872 (byte 109 + 0)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | uint64: 18446744073709551615           | bitfloat32: NaN          | string: -
				"m_bWorldTreeState.0000": uint64(18446744073709551615),
				// manta.(*reader).dumpBits: @ bit 00952 (byte 119 + 0)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | uint64: 18446744073709551615           | bitfloat32: NaN          | string: -
				"m_bWorldTreeState.0001": uint64(18446744073709551615),
				// manta.(*reader).dumpBits: @ bit 11032 (byte 1379 + 0)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | uint64: 18446744073709551615           | bitfloat32: NaN          | string: -
				"m_bWorldTreeState.0126": uint64(18446744073709551615),
				"m_bWorldTreeState.0127": uint64(18446744073709551615),
				// That was 1270 bytes (127 entries) worth of 10-byte varint 18446744073709551615's.
				// Now we see 127 more bytes (127 entries) worth of 1-byte varint 0's.
				// manta.(*reader).dumpBits: @ bit 11112 (byte 1389 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: 0                              | bitfloat32: 0            | string: -
				"m_bWorldTreeState.0128": uint64(0),
				"m_bWorldTreeState.0129": uint64(0),
				// manta.(*reader).dumpBits: @ bit 12128 (byte 1516 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | uint64: ERR                            | bitfloat32: ERR          | string: -
				"m_bWorldTreeState.0255": uint64(0),
			},
		},

		{
			tableName:   "CDOTA_DataRadiant",
			run:         true,
			debug:       false,
			expectCount: (10 + 10 + 1 + 30 + 256),
			expectKeys: map[string]interface{}{
				"m_iReliableGold.0000":      int32(0),
				"m_iReliableGold.0009":      int32(0),
				"m_iUnreliableGold.0000":    int32(625),
				"m_iUnreliableGold.0009":    int32(625),
				"m_iTeamNum":                uint64(2),
				"m_iStartingPositions.0000": int32(-1),
				"m_iStartingPositions.0029": int32(-1),
				"m_bWorldTreeState.0000":    uint64(18446744073709551615),
				"m_bWorldTreeState.0127":    uint64(18446744073709551615),
				"m_bWorldTreeState.0128":    uint64(0),
				"m_bWorldTreeState.0255":    uint64(0),
			},
		},

		{
			tableName:   "CDOTA_DataSpectator",
			run:         true,
			debug:       false,
			expectCount: 12,
			expectKeys: map[string]interface{}{
				// manta.(*reader).dumpBits: @ bit 00021 (byte 002 + 5)  | binary: 1 | uint8: 255 | var32: -8388608    | varu32: 16777215   | varu64: 16777215             | float32: 3.8518597e-34      | float32coord: -16384.969         | string: ERR
				"m_hPrimaryRune": HANDLE_NONE,
				// manta.(*reader).dumpBits: @ bit 00053 (byte 006 + 5)  | binary: 1 | uint8: 255 | var32: -8388608    | varu32: 16777215   | varu64: 16777215             | float32: 3.8518597e-34      | float32coord: -16384.969         | string: ERR
				"m_hSecondaryRune": HANDLE_NONE,
				// manta.(*reader).dumpBits: @ bit 00085 (byte 010 + 5)  | binary: 0 | uint8: 226 | var32: 625         | varu32: 1250       | varu64: 1250                 | float32: 5.4416815e-33      | float32coord: 0.875              | string: ERR
				"m_iNetWorth.0000": int32(625),
				// manta.(*reader).dumpBits: @ bit 00229 (byte 028 + 5)  | binary: 0 | uint8: 226 | var32: 625         | varu32: 1250       | varu64: 1250                 | float32: ERR                | float32coord: 0.875              | string: ERR
				"m_iNetWorth.0009": int32(625),
			},
		},

		{
			tableName:   "CDOTA_DataCustomTeam",
			run:         true,
			debug:       false,
			expectCount: (10 + 10 + 1 + 30 + 256),
			expectKeys: map[string]interface{}{
				"m_iReliableGold.0000":      int32(0),
				"m_iReliableGold.0009":      int32(0),
				"m_iUnreliableGold.0000":    int32(625),
				"m_iUnreliableGold.0009":    int32(625),
				"m_iTeamNum":                uint8(6),
				"m_iStartingPositions.0000": int32(-1),
				"m_iStartingPositions.0029": int32(-1),
				"m_bWorldTreeState.0000":    uint64(18446744073709551615),
				"m_bWorldTreeState.0127":    uint64(18446744073709551615),
				"m_bWorldTreeState.0128":    uint64(0),
				"m_bWorldTreeState.0255":    uint64(0),
			},
		},

		/*
		   class id=413 name=CDOTA_PlayerResource
		   table id=422 name=CDOTA_PlayerResource version=0
		    -> prop 0: type:int32[10](1164) name:m_iTotalEarnedGold(1208) sendNode: (root)(734)
		    -> prop 1: type:int32[10](1164) name:m_iReliableGoldRadiant(1209) sendNode: (root)(734)
		    -> prop 2: type:int32[10](1164) name:m_iReliableGoldDire(1210) sendNode: (root)(734)
		    -> prop 3: type:int32[10](1164) name:m_iUnreliableGoldRadiant(1211) sendNode: (root)(734)
		    -> prop 4: type:int32[10](1164) name:m_iUnreliableGoldDire(1212) sendNode: (root)(734)
		    -> prop 5: type:CUtlSymbolLarge[64](1213) name:m_iszPlayerNames(1214) sendNode: (root)(734)
		    -> prop 6: type:int32[10](1164) name:m_nSelectedHeroID(1215) sendNode: (root)(734)
		    -> prop 7: type:int32[64](1216) name:m_iPlayerTeams(1217) sendNode: (root)(734)
		    -> prop 8: type:int32[64](1216) name:m_iLobbyPlayerTeams(1218) sendNode: (root)(734)
		    -> prop 9: type:int32[64](1216) name:m_iCustomTeamAssignments(1219) sendNode: (root)(734)
		    -> prop 10: type:int32[10](1164) name:m_iKills(1220) sendNode: (root)(734)
		    -> prop 11: type:int32[10](1164) name:m_iAssists(1221) sendNode: (root)(734)
		    -> prop 12: type:int32[10](1164) name:m_iDeaths(1222) sendNode: (root)(734)
		    -> prop 13: type:int32[10](1164) name:m_iStreak(1223) sendNode: (root)(734)
		    -> prop 14: type:int32[10](1164) name:m_iSharedGold(1224) sendNode: (root)(734)
		    -> prop 15: type:int32[10](1164) name:m_iHeroKillGold(1225) sendNode: (root)(734)
		    -> prop 16: type:int32[10](1164) name:m_iCreepKillGold(1226) sendNode: (root)(734)
		    -> prop 17: type:int32[10](1164) name:m_iIncomeGold(1227) sendNode: (root)(734)
		    -> prop 18: type:int32[10](1164) name:m_iLevel(1047) sendNode: (root)(734)
		    -> prop 19: type:int32[10](1164) name:m_iRespawnSeconds(1228) sendNode: (root)(734)
		    -> prop 20: type:int32[10](1164) name:m_iLastBuybackTime(1229) sendNode: (root)(734)
		    -> prop 21: type:int32[10](1164) name:m_iDenyCount(1230) sendNode: (root)(734)
		    -> prop 22: type:int32[10](1164) name:m_iLastHitCount(1231) sendNode: (root)(734)
		    -> prop 23: type:int32[10](1164) name:m_iLastHitStreak(1232) sendNode: (root)(734)
		    -> prop 24: type:int32[10](1164) name:m_iLastHitMultikill(1233) sendNode: (root)(734)
		    -> prop 25: type:int32[10](1164) name:m_iNearbyCreepDeathCount(1234) sendNode: (root)(734)
		    -> prop 26: type:int32[10](1164) name:m_iClaimedDenyCount(1235) sendNode: (root)(734)
		    -> prop 27: type:int32[10](1164) name:m_iClaimedMissCount(1236) sendNode: (root)(734)
		    -> prop 28: type:int32[10](1164) name:m_iMissCount(1237) sendNode: (root)(734)
		    -> prop 29: type:CHandle< CBaseEntity >[10](1004) name:m_hSelectedHero(1238) sendNode: (root)(734)
		    -> prop 30: type:bool[64](1239) name:m_bFullyJoinedServer(1240) sendNode: (root)(734)
		    -> prop 31: type:bool[64](1239) name:m_bFakeClient(1241) sendNode: (root)(734)
		    -> prop 32: type:bool[64](1239) name:m_bIsBroadcaster(1242) sendNode: (root)(734)
		    -> prop 33: type:uint32[64](1243) name:m_iBroadcasterChannel(1244) sendNode: (root)(734)
		    -> prop 34: type:uint32[64](1243) name:m_iBroadcasterChannelSlot(1245) sendNode: (root)(734)
		    -> prop 35: type:bool[64](1239) name:m_bIsBroadcasterChannelCameraman(1246) sendNode: (root)(734)
		    -> prop 36: type:CUtlSymbolLarge[6](1247) name:m_iszBroadcasterChannelDescription(1248) sendNode: (root)(734)
		    -> prop 37: type:CUtlSymbolLarge[6](1247) name:m_iszBroadcasterChannelCountryCode(1249) sendNode: (root)(734)
		    -> prop 38: type:CUtlSymbolLarge[6](1247) name:m_iszBroadcasterChannelLanguageCode(1250) sendNode: (root)(734)
		    -> prop 39: type:int32[64](1216) name:m_iConnectionState(1251) sendNode: (root)(734)
		    -> prop 40: type:bool[10](1252) name:m_bAFK(1253) sendNode: (root)(734)
		    -> prop 41: type:int32[10](1164) name:m_nPossibleHeroSelection(1254) sendNode: (root)(734)
		    -> prop 42: type:int32[20](1255) name:m_nSuggestedHeroes(1256) sendNode: (root)(734)
		    -> prop 43: type:bool[10](1252) name:m_bVoiceChatBanned(1257) sendNode: (root)(734)
		    -> prop 44: type:uint64[64](1258) name:m_iPlayerSteamIDs(1259) sendNode: (root)(734)
		    -> prop 45: type:int32[10](1164) name:m_iTimedRewardDrops(1260) sendNode: (root)(734)
		    -> prop 46: type:int32[10](1164) name:m_iTimedRewardDropOrigins(1261) sendNode: (root)(734)
		    -> prop 47: type:int32[10](1164) name:m_iTimedRewardCrates(1262) sendNode: (root)(734)
		    -> prop 48: type:int32[10](1164) name:m_iTimedRewardEvents(1263) sendNode: (root)(734)
		    -> prop 49: type:uint16[10](1264) name:m_iMetaLevel(1265) sendNode: (root)(734)
		    -> prop 50: type:uint16[10](1264) name:m_iMetaExperience(1266) sendNode: (root)(734)
		    -> prop 51: type:uint16[10](1264) name:m_iMetaExperienceAwarded(1267) sendNode: (root)(734)
		    -> prop 52: type:uint32[10](1268) name:m_iEventPoints(1269) sendNode: (root)(734)
		    -> prop 53: type:uint32[10](1268) name:m_iEventPremiumPoints(1270) sendNode: (root)(734)
		    -> prop 54: type:uint16[10](1264) name:m_iEventRanks(1271) sendNode: (root)(734)
		    -> prop 55: type:uint16[10](1264) name:m_unCompendiumLevel(1272) sendNode: (root)(734)
		    -> prop 56: type:bool[10](1252) name:m_bHasRepicked(1273) sendNode: (root)(734)
		    -> prop 57: type:bool[10](1252) name:m_bHasRandomed(1274) sendNode: (root)(734)
		    -> prop 58: type:bool[10](1252) name:m_bBattleBonusActive(1275) sendNode: (root)(734)
		    -> prop 59: type:uint16[10](1264) name:m_iBattleBonusRate(1276) sendNode: (root)(734)
		    -> prop 60: type:float32[10](1277) name:m_flBuybackCooldownTime(1278) sendNode: (root)(734) bitCount:32
		    -> prop 61: type:float32[10](1277) name:m_flBuybackGoldLimitTime(1279) sendNode: (root)(734) bitCount:32
		    -> prop 62: type:float32[10](1277) name:m_flBuybackCostTime(1280) sendNode: (root)(734) bitCount:32
		    -> prop 63: type:int32[10](1164) name:m_iCustomBuybackCost(1281) sendNode: (root)(734)
		    -> prop 64: type:float32[10](1277) name:m_flCustomBuybackCooldown(1282) sendNode: (root)(734) bitCount:32
		    -> prop 65: type:int32[10](1164) name:m_iGoldBagsCollected(1283) sendNode: (root)(734)
		    -> prop 66: type:float32[10](1277) name:m_fStuns(1284) sendNode: (root)(734) bitCount:32
		    -> prop 67: type:float32[10](1277) name:m_fHealing(1285) sendNode: (root)(734) bitCount:32
		    -> prop 68: type:int32[10](1164) name:m_iTowerKills(1286) sendNode: (root)(734)
		    -> prop 69: type:int32[10](1164) name:m_iRoshanKills(1287) sendNode: (root)(734)
		    -> prop 70: type:CHandle< CBaseEntity >[10](1004) name:m_hCameraTarget(1288) sendNode: (root)(734)
		    -> prop 71: type:Color[10](1289) name:m_CustomPlayerColors(1290) sendNode: (root)(734)
		    -> prop 72: type:uint64[256](1169) name:m_bWorldTreeStateRadiant(1291) sendNode: (root)(734)
		    -> prop 73: type:uint64[256](1169) name:m_bWorldTreeStateDire(1292) sendNode: (root)(734)
		    -> prop 74: type:uint64[128](1293) name:m_bWorldTreeStateSpectator(1294) sendNode: (root)(734)
		    -> prop 75: type:bool[10](1252) name:m_bHasPredictedVictory(1295) sendNode: (root)(734)
		    -> prop 76: type:bool[10](1252) name:m_bReservedHeroOnly(1296) sendNode: (root)(734)
		    -> prop 77: type:bool[10](1252) name:m_bQualifiesForPAContractReward(1297) sendNode: (root)(734)
		    -> prop 78: type:int32[10](1164) name:m_UnitShareMasks(1298) sendNode: (root)(734)
		    -> prop 79: type:int32[10](1164) name:m_iTotalEarnedXP(1299) sendNode: (root)(734)
		*/
		{
			tableName:   "CDOTA_PlayerResource",
			run:         true,
			debug:       false,
			expectCount: 2056,
			expectKeys: map[string]interface{}{
				// manta.(*reader).dumpBits: @ bit 04274 (byte 534 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iTotalEarnedGold.0000": int32(0),
				// ...
				// manta.(*reader).dumpBits: @ bit 04346 (byte 543 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iTotalEarnedGold.0009": int32(0),
				// manta.(*reader).dumpBits: @ bit 04354 (byte 544 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iReliableGoldRadiant.0000": int32(0),
				// ...
				// manta.(*reader).dumpBits: @ bit 04426 (byte 553 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iReliableGoldRadiant.0009": int32(0),
				// manta.(*reader).dumpBits: @ bit 04434 (byte 554 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iReliableGoldDire.0000": int32(0),
				// ...
				// manta.(*reader).dumpBits: @ bit 04507 (byte 563 + 3)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iReliableGoldDire.0009": int32(0),
				// manta.(*reader).dumpBits: @ bit 04514 (byte 564 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iUnreliableGoldRadiant.0000": int32(0),
				// ...
				// manta.(*reader).dumpBits: @ bit 04586 (byte 573 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iUnreliableGoldRadiant.0009": int32(0),
				// manta.(*reader).dumpBits: @ bit 04594 (byte 574 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iUnreliableGoldDire.0000": int32(0),
				// ...
				// manta.(*reader).dumpBits: @ bit 04667 (byte 583 + 3)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: -3.3558172e-07 | string: -
				"m_iUnreliableGoldDire.0009": int32(0),
				// manta.(*reader).dumpBits: @ bit 04674 (byte 584 + 2)  | binary: 0 | uint8: 84  | var32: 42          | varu32: 84         | varu64: 84                   | float32: 2.6910056e+20 | string: Thias
				"m_iszPlayerNames.0000": "Thias",
				// manta.(*reader).dumpBits: @ bit 04722 (byte 590 + 2)  | binary: 1 | uint8: 65  | var32: -33         | varu32: 65         | varu64: 65                   | float32: 9.592237e-39 | string: Ash
				"m_iszPlayerNames.0001": "Ash",
				// manta.(*reader).dumpBits: @ bit 04754 (byte 594 + 2)  | binary: 1 | uint8: 107 | var32: -54         | varu32: 107        | varu64: 107                  | float32: 7.007159e+22 | string: kimee
				"m_iszPlayerNames.0002": "kimee",
				// manta.(*reader).dumpBits: @ bit 04802 (byte 600 + 2)  | binary: 0 | uint8: 68  | var32: 34          | varu32: 68         | varu64: 68                   | float32: 2.0531703e-19 | string: Der Baerenjurek
				"m_iszPlayerNames.0003": "Der Baerenjurek",
				// manta.(*reader).dumpBits: @ bit 04930 (byte 616 + 2)  | binary: 1 | uint8: 109 | var32: -55         | varu32: 109        | varu64: 109                  | float32: 6.947208e+22 | string: makes
				"m_iszPlayerNames.0004": "makes",
				// manta.(*reader).dumpBits: @ bit 04978 (byte 622 + 2)  | binary: 0 | uint8: 208 | var32: 64619560    | varu32: 129239120  | varu64: ERR                  | float32: -0.10183871  | string: -
				"m_iszPlayerNames.0005": "Анонимный геймер",
				// manta.(*reader).dumpBits: @ bit 05234 (byte 654 + 2)  | binary: 1 | uint8: 97  | var32: -49         | varu32: 97         | varu64: 97                   | float32: 6.890133e+22 | string: ariethebeast
				"m_iszPlayerNames.0006": "ariethebeast",
				// manta.(*reader).dumpBits: @ bit 05338 (byte 667 + 2)  | binary: 1 | uint8: 79  | var32: -40         | varu32: 79         | varu64: 79                   | float32: 1.7408505e+25 | string: Officer Pacman
				"m_iszPlayerNames.0007": "Officer Pacman",
				// manta.(*reader).dumpBits: @ bit 05458 (byte 682 + 2)  | binary: 1 | uint8: 91  | var32: -46         | varu32: 91         | varu64: 91                   | float32: 9.562119e+17 | string: -
				"m_iszPlayerNames.0008": "[RT]FleeX",
				"m_iszPlayerNames.0009": "[RT]G@dget",
				"m_iszPlayerNames.0062": "",
				"m_iszPlayerNames.0063": "",

				// manta.(*reader).dumpBits: @ bit 06058 (byte 757 + 2)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | float32: 2.3694278e-38 | string: -
				"m_nSelectedHeroID.0000": int32(-1),
				// ....
				// manta.(*reader).dumpBits: @ bit 06130 (byte 766 + 2)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | float32: 1.5518364e-36 | string: -
				"m_nSelectedHeroID.0009": int32(-1),

				// manta.(*reader).dumpBits: @ bit 06138 (byte 767 + 2)  | binary: 0 | uint8: 4   | var32: 2           | varu32: 4          | varu64: 4                    | float32: 1.551837e-36 | string: -
				"m_iPlayerTeams.0000": int32(2),
				// manta.(*reader).dumpBits: @ bit 06146 (byte 768 + 2)  | binary: 0 | uint8: 4   | var32: 2           | varu32: 4          | varu64: 4                    | float32: 1.551837e-36 | string: -
				"m_iPlayerTeams.0001": int32(2),
				// manta.(*reader).dumpBits: @ bit 06154 (byte 769 + 2)  | binary: 0 | uint8: 4   | var32: 2           | varu32: 4          | varu64: 4                    | float32: 2.482939e-35 | string: -
				"m_iPlayerTeams.0002": int32(2),
				// manta.(*reader).dumpBits: @ bit 06162 (byte 770 + 2)  | binary: 0 | uint8: 4   | var32: 2           | varu32: 4          | varu64: 4                    | float32: 2.520555e-35 | string: -
				"m_iPlayerTeams.0003": int32(2),
				// manta.(*reader).dumpBits: @ bit 06170 (byte 771 + 2)  | binary: 0 | uint8: 4   | var32: 2           | varu32: 4          | varu64: 4                    | float32: 2.5207018e-35 | string: -
				"m_iPlayerTeams.0004": int32(2),
				// manta.(*reader).dumpBits: @ bit 06178 (byte 772 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 2.5207024e-35 | string: -
				"m_iPlayerTeams.0005": int32(3),
				// manta.(*reader).dumpBits: @ bit 06186 (byte 773 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 2.5207024e-35 | string: -
				"m_iPlayerTeams.0006": int32(3),
				// manta.(*reader).dumpBits: @ bit 06194 (byte 774 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 6.452998e-33 | string: -
				"m_iPlayerTeams.0007": int32(3),
				// manta.(*reader).dumpBits: @ bit 06202 (byte 775 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 6.645591e-33 | string: -
				"m_iPlayerTeams.0008": int32(3),
				// manta.(*reader).dumpBits: @ bit 06210 (byte 776 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 6.6463435e-33 | string: -
				"m_iPlayerTeams.0009": int32(3),
				// manta.(*reader).dumpBits: @ bit 06218 (byte 777 + 2)  | binary: 0 | uint8: 10  | var32: 5           | varu32: 10         | varu64: 10                   | float32: 6.6463464e-33 | string: -
				"m_iPlayerTeams.0010": int32(5), // spectators
				// manta.(*reader).dumpBits: @ bit 06642 (byte 830 + 2)  | binary: 0 | uint8: 10  | var32: 5           | varu32: 10         | varu64: 10                   | float32: 1.551838e-36 | string: -
				"m_iPlayerTeams.0063": int32(5), // spectators

				// manta.(*reader).dumpBits: @ bit 06650 (byte 831 + 2)  | binary: 0 | uint8: 4   | var32: 2           | varu32: 4          | varu64: 4                    | float32: 1.551837e-36 | string: -
				"m_iLobbyPlayerTeams.0000": int32(2),
				"m_iLobbyPlayerTeams.0001": int32(2),
				"m_iLobbyPlayerTeams.0002": int32(2),
				"m_iLobbyPlayerTeams.0003": int32(2),
				"m_iLobbyPlayerTeams.0004": int32(2),
				// manta.(*reader).dumpBits: @ bit 06690 (byte 836 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 2.5207024e-35 | string: -
				"m_iLobbyPlayerTeams.0005": int32(3),
				"m_iLobbyPlayerTeams.0006": int32(3),
				"m_iLobbyPlayerTeams.0007": int32(3),
				"m_iLobbyPlayerTeams.0008": int32(3),
				"m_iLobbyPlayerTeams.0009": int32(3),
				// ...
				// manta.(*reader).dumpBits: @ bit 07154 (byte 894 + 2)  | binary: 0 | uint8: 10  | var32: 5           | varu32: 10         | varu64: 10                   | float32: 1.4e-44      | string:
				"m_iLobbyPlayerTeams.0063": int32(5),
				// manta.(*reader).dumpBits: @ bit 07162 (byte 895 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iCustomTeamAssignments.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 07234 (byte 904 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 6.646339e-33 | string: -
				"m_iCustomTeamAssignments.0009": int32(0),
				// manta.(*reader).dumpBits: @ bit 07242 (byte 905 + 2)  | binary: 0 | uint8: 10  | var32: 5           | varu32: 10         | varu64: 10                   | float32: 6.6463464e-33 | string:
				"m_iCustomTeamAssignments.0010": int32(5),
				// ...
				// manta.(*reader).dumpBits: @ bit 07666 (byte 958 + 2)  | binary: 0 | uint8: 10  | var32: 5           | varu32: 10         | varu64: 10                   | float32: 1.4e-44      | string:
				"m_iCustomTeamAssignments.0063": int32(5),

				// manta.(*reader).dumpBits: @ bit 07674 (byte 959 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iKills.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 07754 (byte 969 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iAssists.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 07834 (byte 979 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iDeaths.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 07914 (byte 989 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iStreak.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 07994 (byte 999 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iSharedGold.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08074 (byte 1009 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iHeroKillGold.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08154 (byte 1019 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iCreepKillGold.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08234 (byte 1029 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iIncomeGold.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08314 (byte 1039 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iLevel.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08394 (byte 1049 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iRespawnSeconds.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08474 (byte 1059 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iLastBuybackTime.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08554 (byte 1069 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iDenyCount.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08634 (byte 1079 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iLastHitCount.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08714 (byte 1089 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iLastHitStreak.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08794 (byte 1099 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iLastHitMultikill.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08874 (byte 1109 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iNearbyCreepDeathCount.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 08954 (byte 1119 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iClaimedDenyCount.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 09034 (byte 1129 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iClaimedMissCount.0000": int32(0),
				// manta.(*reader).dumpBits: @ bit 09114 (byte 1139 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iMissCount.0000": int32(0),

				"m_bIsBroadcaster.0000": false,
				// ... all the same ...
				// manta.(*reader).dumpBits: @ bit 09705 (byte 1213 + 1)  | binary: 0 | uint8: 12  | var32: 6           | varu32: 12         | varu64: 12                   | float32: 1.0788833e-31 | string:
				"m_bIsBroadcaster.0063": false,

				// manta.(*reader).dumpBits: @ bit 09706 (byte 1213 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 2.5207024e-35 | string: -
				"m_iBroadcasterChannel.0000": uint64(6),
				// ... all the same ...
				// manta.(*reader).dumpBits: @ bit 10210 (byte 1276 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 2.5207024e-35 | string: -
				"m_iBroadcasterChannel.0063": uint64(6),

				// manta.(*reader).dumpBits: @ bit 10218 (byte 1277 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 2.5207024e-35 | string: -
				"m_iBroadcasterChannelSlot.0000": uint64(6),
				// ... all the same ...
				// manta.(*reader).dumpBits: @ bit 10722 (byte 1340 + 2)  | binary: 0 | uint8: 6   | var32: 3           | varu32: 6          | varu64: 6                    | float32: 8e-45        | string: -
				"m_iBroadcasterChannelSlot.0063": uint64(6),

				// manta.(*reader).dumpBits: @ bit 10730 (byte 1341 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bIsBroadcasterChannelCameraman.0000": false,
				// ... all the same ...
				// manta.(*reader).dumpBits: @ bit 10793 (byte 1349 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bIsBroadcasterChannelCameraman.0063": false,

				// manta.(*reader).dumpBits: @ bit 10794 (byte 1349 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iszBroadcasterChannelDescription.0000": "",
				// manta.(*reader).dumpBits: @ bit 10834 (byte 1354 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iszBroadcasterChannelDescription.0005": "",

				// manta.(*reader).dumpBits: @ bit 10842 (byte 1355 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iszBroadcasterChannelCountryCode.0000": "",
				// manta.(*reader).dumpBits: @ bit 10882 (byte 1360 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iszBroadcasterChannelCountryCode.0005": "",

				// manta.(*reader).dumpBits: @ bit 10890 (byte 1361 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iszBroadcasterChannelLanguageCode.0000": "",
				// manta.(*reader).dumpBits: @ bit 10930 (byte 1366 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 9.551466e-38 | string: -
				"m_iszBroadcasterChannelLanguageCode.0005": "",

				// manta.(*reader).dumpBits: @ bit 10938 (byte 1367 + 2)  | binary: 0 | uint8: 2   | var32: 1           | varu32: 2          | varu64: 2                    | float32: 9.551468e-38 | string: -
				"m_iConnectionState.0000": int32(1),
				// all the same
				// manta.(*reader).dumpBits: @ bit 11442 (byte 1430 + 2)  | binary: 0 | uint8: 2   | var32: 1           | varu32: 2          | varu64: 2                    | float32: 1.5516529e-36 | string: -
				"m_iConnectionState.0063": int32(1),

				// manta.(*reader).dumpBits: @ bit 11450 (byte 1431 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 1.5518362e-36 | string: -
				"m_bAFK.0000": false,
				// manta.(*reader).dumpBits: @ bit 11459 (byte 1432 + 3)  | binary: 0 | uint8: 2   | var32: 1           | varu32: 2          | varu64: 2                    | float32: 9.551468e-38 | string: -
				"m_bAFK.0009": false,

				// manta.(*reader).dumpBits: @ bit 11460 (byte 1432 + 4)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | float32: 2.3694278e-38 | string: -
				"m_nPossibleHeroSelection.0000": int32(-1),

				// manta.(*reader).dumpBits: @ bit 11540 (byte 1442 + 4)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | float32: 2.3694278e-38 | string: -
				"m_nSuggestedHeroes.0000": int32(-1),
				// ... all the same ...
				// manta.(*reader).dumpBits: @ bit 11692 (byte 1461 + 4)  | binary: 1 | uint8: 1   | var32: -1          | varu32: 1          | varu64: 1                    | float32: -0.22265626  | string: -
				"m_nSuggestedHeroes.0019": int32(-1),

				// manta.(*reader).dumpBits: @ bit 11700 (byte 1462 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 2.0642193e-17 | string: -
				"m_bVoiceChatBanned.0000": false,
				"m_bVoiceChatBanned.0009": false,

				// manta.(*reader).dumpBits: @ bit 11710 (byte 1463 + 6)  | binary: 1 | uint8: 153 | var32: -17898445   | varu32: 35796889   | varu64: 76561197996062617    | float32: -2.160468e-28 | string: -
				"m_iPlayerSteamIDs.0000": uint64(76561197996062617),
				"m_iPlayerSteamIDs.0001": uint64(76561198046993147),
				"m_iPlayerSteamIDs.0002": uint64(76561197961237397),
				"m_iPlayerSteamIDs.0003": uint64(76561197973633834),
				"m_iPlayerSteamIDs.0004": uint64(76561197977509327),
				"m_iPlayerSteamIDs.0005": uint64(76561198122536495),
				"m_iPlayerSteamIDs.0006": uint64(76561198065323776),
				"m_iPlayerSteamIDs.0007": uint64(76561198054698773),
				"m_iPlayerSteamIDs.0008": uint64(76561198134851521),
				// manta.(*reader).dumpBits: @ bit 12358 (byte 1544 + 6)  | binary: 0 | uint8: 148 | var32: 85634058    | varu32: 171268116  | varu64: 76561198131533844    | float32: -1.14723815e+11 | string: -
				"m_iPlayerSteamIDs.0009": uint64(76561198131533844),
				// manta.(*reader).dumpBits: @ bit 12430 (byte 1553 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iPlayerSteamIDs.0010": uint64(0),
				// ...
				// manta.(*reader).dumpBits: @ bit 12854 (byte 1606 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iPlayerSteamIDs.0063": uint64(0),

				// manta.(*reader).dumpBits: @ bit 12862 (byte 1607 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iTimedRewardDrops.0000": int32(0),

				// manta.(*reader).dumpBits: @ bit 12942 (byte 1617 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iTimedRewardDropOrigins.0000": int32(0),

				// manta.(*reader).dumpBits: @ bit 13022 (byte 1627 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iTimedRewardCrates.0000": int32(0),

				// manta.(*reader).dumpBits: @ bit 13110 (byte 1638 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iTimedRewardEvents.0000": int32(0),

				// manta.(*reader).dumpBits: @ bit 13182 (byte 1647 + 6)  | binary: 0 | uint8: 66  | var32: 33          | varu32: 66         | varu64: 66                   | float32: 2.5220848e-32 | string: -
				"m_iMetaLevel.0000": uint64(66), // Dota profile level
				// manta.(*reader).dumpBits: @ bit 13190 (byte 1648 + 6)  | binary: 0 | uint8: 244 | var32: 186         | varu32: 372        | varu64: 372                  | float32: -1.6340727e-36 | string: -
				"m_iMetaLevel.0001": uint64(372),
				// manta.(*reader).dumpBits: @ bit 13206 (byte 1650 + 6)  | binary: 1 | uint8: 11  | var32: -6          | varu32: 11         | varu64: 11                   | float32: 2.4359213e-35 | string: -
				"m_iMetaLevel.0002": uint64(11),
				"m_iMetaLevel.0003": uint64(132),
				"m_iMetaLevel.0004": uint64(6),
				"m_iMetaLevel.0005": uint64(53),
				"m_iMetaLevel.0006": uint64(18),
				"m_iMetaLevel.0007": uint64(164),
				"m_iMetaLevel.0008": uint64(88),
				"m_iMetaLevel.0009": uint64(46),

				// manta.(*reader).dumpBits: @ bit 13286 (byte 1660 + 6)  | binary: 0 | uint8: 188 | var32: 94          | varu32: 188        | varu64: 188                  | float32: 1.2601937e-35 | string: -
				"m_iMetaExperience.0000": uint64(188),
				"m_iMetaExperience.0001": uint64(646),
				"m_iMetaExperience.0002": uint64(970),
				"m_iMetaExperience.0003": uint64(529),
				"m_iMetaExperience.0004": uint64(268),
				"m_iMetaExperience.0005": uint64(81),
				"m_iMetaExperience.0006": uint64(114),
				"m_iMetaExperience.0007": uint64(511),
				"m_iMetaExperience.0008": uint64(232),
				"m_iMetaExperience.0009": uint64(342),

				// manta.(*reader).dumpBits: @ bit 13430 (byte 1678 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iMetaExperienceAwarded.0000": uint64(0),
				"m_iMetaExperienceAwarded.0001": uint64(0),
				"m_iMetaExperienceAwarded.0002": uint64(0),
				"m_iMetaExperienceAwarded.0003": uint64(0),
				"m_iMetaExperienceAwarded.0004": uint64(0),
				"m_iMetaExperienceAwarded.0005": uint64(0),
				"m_iMetaExperienceAwarded.0006": uint64(0),
				"m_iMetaExperienceAwarded.0007": uint64(0),
				"m_iMetaExperienceAwarded.0008": uint64(0),
				"m_iMetaExperienceAwarded.0009": uint64(0),

				// manta.(*reader).dumpBits: @ bit 13510 (byte 1688 + 6)  | binary: 0 | uint8: 212 | var32: 14314       | varu32: 28628      | varu64: 28628                | float32: -2.5613195e-29 | string: -
				"m_iEventPoints.0000": uint64(28628), // compendium level 286
				"m_iEventPoints.0001": uint64(3600),  // compendium level 36
				"m_iEventPoints.0002": uint64(0),     // no compendium
				"m_iEventPoints.0003": uint64(8881),  // compendium level 89
				"m_iEventPoints.0004": uint64(0),
				"m_iEventPoints.0005": uint64(0),
				"m_iEventPoints.0006": uint64(0),
				"m_iEventPoints.0007": uint64(0),
				"m_iEventPoints.0008": uint64(0),    // no compendium
				"m_iEventPoints.0009": uint64(7283), // compendium, unknown level

				// manta.(*reader).dumpBits: @ bit 13630 (byte 1703 + 6)  | binary: 1 | uint8: 21  | var32: -11         | varu32: 21         | varu64: 21                   | float32: 3.994874e-39 | string: -
				"m_iEventPremiumPoints.0000": uint64(21),
				"m_iEventPremiumPoints.0001": uint64(5504),
				"m_iEventPremiumPoints.0002": uint64(0),
				"m_iEventPremiumPoints.0003": uint64(294),
				"m_iEventPremiumPoints.0004": uint64(46),
				"m_iEventPremiumPoints.0005": uint64(5),
				"m_iEventPremiumPoints.0006": uint64(0),
				"m_iEventPremiumPoints.0007": uint64(395),
				"m_iEventPremiumPoints.0008": uint64(70),
				"m_iEventPremiumPoints.0009": uint64(471),

				// manta.(*reader).dumpBits: @ bit 13742 (byte 1717 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iEventRanks.0000": uint64(0),
				"m_iEventRanks.0001": uint64(0),
				"m_iEventRanks.0002": uint64(0),
				"m_iEventRanks.0003": uint64(0),
				"m_iEventRanks.0004": uint64(0),
				"m_iEventRanks.0005": uint64(0),
				"m_iEventRanks.0006": uint64(0),
				"m_iEventRanks.0007": uint64(0),
				"m_iEventRanks.0008": uint64(0),
				"m_iEventRanks.0009": uint64(0),

				// manta.(*reader).dumpBits: @ bit 13822 (byte 1727 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_unCompendiumLevel.0000": uint64(0),
				"m_unCompendiumLevel.0001": uint64(0),
				"m_unCompendiumLevel.0002": uint64(0),
				"m_unCompendiumLevel.0003": uint64(0),
				"m_unCompendiumLevel.0004": uint64(0),
				"m_unCompendiumLevel.0005": uint64(0),
				"m_unCompendiumLevel.0006": uint64(0),
				"m_unCompendiumLevel.0007": uint64(0),
				"m_unCompendiumLevel.0008": uint64(0),
				"m_unCompendiumLevel.0009": uint64(0),

				// manta.(*reader).dumpBits: @ bit 13902 (byte 1737 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bHasRepicked.0000": false,
				"m_bHasRepicked.0009": false,

				// manta.(*reader).dumpBits: @ bit 13913 (byte 1739 + 1)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bHasRandomed.0000": false,
				"m_bHasRandomed.0009": false,

				// manta.(*reader).dumpBits: @ bit 13922 (byte 1740 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bBattleBonusActive.0000": false,
				"m_bBattleBonusActive.0009": false,

				// manta.(*reader).dumpBits: @ bit 13932 (byte 1741 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iBattleBonusRate.0000": uint64(0),
				"m_iBattleBonusRate.0009": uint64(0),

				// !! read 32 bits for each float. it has a bitCount of 32.
				// manta.(*reader).dumpBits: @ bit 14012 (byte 1751 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_flBuybackCooldownTime.0000": float32(0),
				"m_flBuybackCooldownTime.0009": float32(0),

				// !!! read 32 bits for each float. it has a bitCount of 32.
				// manta.(*reader).dumpBits: @ bit 14332 (byte 1791 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_flBuybackGoldLimitTime.0000": float32(0),
				"m_flBuybackGoldLimitTime.0009": float32(0),

				// !! read 32 bits for each float. it has a bitCount of 32.
				// manta.(*reader).dumpBits: @ bit 14652 (byte 1831 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_flBuybackCostTime.0000": float32(0),
				"m_flBuybackCostTime.0009": float32(0),

				// manta.(*reader).dumpBits: @ bit 14972 (byte 1871 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iCustomBuybackCost.0000": int32(0),
				"m_iCustomBuybackCost.0009": int32(0),

				// !! read 32 bits for each float. it has a bitCount of 32.
				// manta.(*reader).dumpBits: @ bit 15052 (byte 1881 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_flCustomBuybackCooldown.0000": float32(0),
				"m_flCustomBuybackCooldown.0009": float32(0),

				// manta.(*reader).dumpBits: @ bit 15372 (byte 1921 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iGoldBagsCollected.0000": int32(0),
				"m_iGoldBagsCollected.0009": int32(0),

				// !! read 32 bits for each float. it has a bitCount of 32.
				// manta.(*reader).dumpBits: @ bit 15452 (byte 1931 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_fStuns.0000": float32(0),
				"m_fStuns.0009": float32(0),

				// !! read 32 bits for each float. it has a bitCount of 32.
				// manta.(*reader).dumpBits: @ bit 15772 (byte 1971 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_fHealing.0000": float32(0),
				"m_fHealing.0009": float32(0),

				// manta.(*reader).dumpBits: @ bit 16092 (byte 2011 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iTowerKills.0000": int32(0),
				"m_iTowerKills.0009": int32(0),

				// manta.(*reader).dumpBits: @ bit 16172 (byte 2021 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iRoshanKills.0000": int32(0),
				"m_iRoshanKills.0009": int32(0),

				// manta.readProperties: reading type:CHandle< CBaseEntity >[10](1004) name:m_hCameraTarget(1288) from position 16252/49616
				// manta.readProperties: WARN: reading m_hCameraTarget.0 (CHandle< CBaseEntity >) as varint32
				// manta.(*reader).dumpBits: @ bit 16252 (byte 2031 + 4)  | binary: 1 | uint8: 255 | var32: -8388608    | varu32: 16777215   | varu64: 16777215             | float32: 3.8518597e-34 | string: -
				"m_hCameraTarget.0000": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0001": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0002": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0003": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0004": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0005": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0006": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0007": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0008": HANDLE_NONE, // these align but value looks wrong
				"m_hCameraTarget.0009": HANDLE_NONE, // these align but value looks wrong

				// manta.readProperties: reading type:Color[10](1289) name:m_CustomPlayerColors(1290) from position 16572/49616
				// manta.readProperties: WARN: reading m_CustomPlayerColors.0 (Color) as varint32
				// manta.(*reader).dumpBits: @ bit 16572 (byte 2071 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_CustomPlayerColors.0000": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0001": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0002": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0003": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0004": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0005": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0006": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0007": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0008": int32(0), // each gets read as a 1-byte varint
				"m_CustomPlayerColors.0009": int32(0), // each gets read as a 1-byte varint

				// manta.readProperties: reading type:uint64[256](1169) name:m_bWorldTreeStateRadiant(1291) from position 16652/49616
				// manta.(*reader).dumpBits: @ bit 16652 (byte 2081 + 4)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | float32: NaN          | string: -
				"m_bWorldTreeStateRadiant.0000": uint64(18446744073709551615),
				// ... all the same ...
				"m_bWorldTreeStateRadiant.0127": uint64(18446744073709551615),
				// manta.(*reader).dumpBits: @ bit 26892 (byte 3361 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bWorldTreeStateRadiant.0128": uint64(0),
				// ... all the same ...
				"m_bWorldTreeStateRadiant.0255": uint64(0),

				// manta.(*reader).dumpBits: @ bit 27916 (byte 3489 + 4)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | float32: NaN          | string: -
				"m_bWorldTreeStateDire.0000": uint64(18446744073709551615),
				"m_bWorldTreeStateDire.0127": uint64(18446744073709551615),
				// manta.(*reader).dumpBits: @ bit 38156 (byte 4769 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bWorldTreeStateDire.0128": uint64(0),
				"m_bWorldTreeStateDire.0255": uint64(0),
				// manta.(*reader).dumpBits: @ bit 39180 (byte 4897 + 4)  | binary: 1 | uint8: 255 | var32: -2147483648 | varu32: 4294967295 | varu64: 18446744073709551615 | float32: NaN          | string: -
				"m_bWorldTreeStateSpectator.0000": uint64(18446744073709551615),
				"m_bWorldTreeStateSpectator.0127": uint64(18446744073709551615),

				// manta.(*reader).dumpBits: @ bit 49420 (byte 6177 + 4)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bHasPredictedVictory.0000": false,
				"m_bHasPredictedVictory.0009": false,

				// manta.(*reader).dumpBits: @ bit 49430 (byte 6178 + 6)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bReservedHeroOnly.0000": false,
				"m_bReservedHeroOnly.0009": false,

				// manta.(*reader).dumpBits: @ bit 49440 (byte 6180 + 0)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_bQualifiesForPAContractReward.0000": false,
				"m_bQualifiesForPAContractReward.0009": false,

				// manta.(*reader).dumpBits: @ bit 49450 (byte 6181 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_UnitShareMasks.0000": int32(0),
				"m_UnitShareMasks.0009": int32(0),

				// manta.(*reader).dumpBits: @ bit 49530 (byte 6191 + 2)  | binary: 0 | uint8: 0   | var32: 0           | varu32: 0          | varu64: 0                    | float32: 0            | string: -
				"m_iTotalEarnedXP.0000": int32(0),
				"m_iTotalEarnedXP.0009": int32(0),
			},
		},

		{
			tableName:   "CFogController",
			run:         true,
			debug:       false,
			expectCount: 21,
			expectKeys: map[string]interface{}{
				"colorPrimary":         uint64(4294963904),
				"colorSecondary":       uint64(33554431),
				"colorPrimaryLerpTo":   uint64(4294963904),
				"colorSecondaryLerpTo": uint64(33554431),
				"enable":               true,
				"blend":                false,
				"m_bNoReflectionFog":   true,
			},
		},

		{
			tableName:   "CDOTASpectatorGraphManagerProxy",
			run:         true,
			debug:       false,
			expectCount: 398,
			expectKeys: map[string]interface{}{
				"CDOTASpectatorGraphManager.m_rgPlayerGraphData.0000": uint32(2097151),
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

		/*
			WIP:

			{
				tableName:   "CBaseAnimating",
				run:         false,
				debug:       false,
				expectCount: 0,
				expectKeys:  map[string]interface{}{},
			},
		*/
	}

	// Load our send tables
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables/1560315800.pbmsg"), m); err != nil {
		panic(err)
	}

	// Retrieve the flattened field serializer
	fs := ParseSendTables(m, GetDefaultPropertySerializerTable())

	// Iterate through scenarios
	for _, s := range scenarios {
		// Load up a fixture
		buf := _read_fixture(_sprintf("instancebaseline/1560315800_%s.rawbuf", s.tableName))

		serializer := fs.Serializers[s.tableName][0]
		assert.NotNil(serializer)

		// Optionally skip
		if !s.run {
			continue
		}

		// Optionally disable debugging
		debugMode = s.debug

		// Read properties
		r := NewReader(buf)
		props := ReadProperties(r, serializer)
		assert.Equal(len(props), s.expectCount)

		for k, v := range s.expectKeys {
			assert.EqualValues(v, props[k])
		}

		// There shouldn't be more than 8 bits left in the buffer
		_debugf("Remaining bits %v", r.remBits())
		assert.True(r.remBits() < 8)
	}
}
