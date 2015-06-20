package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestSendPropTypeInfo(t *testing.T) {
	assert := assert.New(t)

	scenarios := []struct {
		in    string
		name  string
		count int
	}{
		{"uint8", "uint8", 1},
		{"CUtlVector< CHandle< CBasePlayer > >", "CUtlVector< CHandle< CBasePlayer > >", 1},
		{"CUtlVector< CHandle< CBasePlayer > >[16]", "CUtlVector< CHandle< CBasePlayer > >", 16},
		{"CEntityIdentity*", "CEntityIdentity*", 1},
		{"uint32[32]", "uint32", 32},
		{"int32", "int32", 1},
	}

	for _, s := range scenarios {
		sp := &sendProp{dtName: s.in}
		name, count, err := sp.typeInfo()
		assert.Equal(s.name, name, s.in)
		assert.Equal(s.count, count, s.in)
		assert.Nil(err, s.in)
	}
}

func TestSendTableParsing(t *testing.T) {
	assert := assert.New(t)

	// The single message from a real match
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables/1560315800.pbmsg"), m); err != nil {
		panic(err)
	}

	// Just a basic length check
	assert.Equal(185229, len(m.GetData()))

	// Parse the send tables
	st, err := parseSendTables(m)
	assert.Nil(err)

	// Verify the tables
	assert.Equal(685, len(st.tables))
	assert.Equal("CDOTA_NPC_Observer_Ward", st.tables["CDOTA_NPC_Observer_Ward"].name)
	assert.Equal("CBaseAnimating", st.tables["CBaseAnimating"].name)

	// Verify the properties
	assert.Equal(1303, len(st.props))

	assert.Equal("uint16", st.props[0].dtName)
	assert.Equal("m_cellX", st.props[0].varName)

	assert.Equal("bool", st.props[200].dtName)
	assert.Equal("m_bAllowAutoMovement", st.props[200].varName)
}
