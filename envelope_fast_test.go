package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

// TestFullPacketData verifies the hand-rolled CDemoFullPacket envelope scan
// extracts the same packet data (and packet presence) as the full protobuf
// unmarshal that onCDemoFullPacket consumes, including protobuf's merge of a
// repeated singular message field.
func TestFullPacketData(t *testing.T) {
	assert := assert.New(t)

	tables := &dota.CDemoStringTables{
		Tables: []*dota.CDemoStringTablesTableT{{
			TableName: proto.String("ActiveModifiers"),
			Items: []*dota.CDemoStringTablesItemsT{
				{Str: proto.String("a"), Data: []byte{1, 2, 3}},
				{Str: proto.String("b")},
			},
		}},
	}

	mustMarshal := func(m proto.Message) []byte {
		buf, err := proto.Marshal(m)
		if err != nil {
			t.Fatal(err)
		}
		return buf
	}
	concat := func(bufs ...[]byte) []byte {
		var out []byte
		for _, b := range bufs {
			out = append(out, b...)
		}
		return out
	}

	scenarios := map[string][]byte{
		"empty":         mustMarshal(&dota.CDemoFullPacket{}),
		"tables only":   mustMarshal(&dota.CDemoFullPacket{StringTable: tables}),
		"empty packet":  mustMarshal(&dota.CDemoFullPacket{Packet: &dota.CDemoPacket{}}),
		"packet only":   mustMarshal(&dota.CDemoFullPacket{Packet: &dota.CDemoPacket{Data: []byte{9, 8, 7}}}),
		"tables+packet": mustMarshal(&dota.CDemoFullPacket{StringTable: tables, Packet: &dota.CDemoPacket{Data: []byte{4, 5}}}),
		// Two serialized messages back to back merge: the later data wins,
		// and a later packet without data keeps the earlier data.
		"merge replaces": concat(
			mustMarshal(&dota.CDemoFullPacket{Packet: &dota.CDemoPacket{Data: []byte{1}}}),
			mustMarshal(&dota.CDemoFullPacket{StringTable: tables, Packet: &dota.CDemoPacket{Data: []byte{2, 3}}}),
		),
		"merge keeps": concat(
			mustMarshal(&dota.CDemoFullPacket{Packet: &dota.CDemoPacket{Data: []byte{1}}}),
			mustMarshal(&dota.CDemoFullPacket{Packet: &dota.CDemoPacket{}}),
		),
	}

	for name, buf := range scenarios {
		want := &dota.CDemoFullPacket{}
		if err := proto.Unmarshal(buf, want); err != nil {
			t.Fatalf("%s: %s", name, err)
		}

		data, hasPacket, err := fullPacketData(buf)
		assert.NoError(err, name)
		assert.Equal(want.Packet != nil, hasPacket, name)
		assert.Equal(want.GetPacket().GetData(), data, name)
	}

	_, _, err := fullPacketData([]byte{0x12, 0x05, 0x1a})
	assert.Error(err, "truncated packet field")
}
