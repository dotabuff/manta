package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestReadProperties(t *testing.T) {
	t.Skip() // Skip for now. Remove this while working on readProperties.

	assert := assert.New(t)

	scenarios := []struct {
		tableName string
		propCount int
	}{
		{"CBaseEntity", 12345},
		{"CDOTA_BaseNPC", 12345},
		{"CDOTA_Unit_Roshan", 12345},
		{"CWorld", 12345},
		{"CDOTA_Item", 12345},
		{"CDOTA_Item_Physical", 12345},
	}

	// Load our send tables
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables/01.pbmsg"), m); err != nil {
		panic(err)
	}

	st, err := parseSendTables(m)
	assert.Nil(err)

	// Iterate through scenarios
	for _, s := range scenarios {
		// Load up a fixture
		buf := _read_fixture(_sprintf("instancebaseline/%s.raw", s.tableName))
		st, ok := st.getTableByName(s.tableName)
		assert.True(ok)

		// Read properties
		r := newReader(buf)
		props := readProperties(r, st)
		assert.Equal(s.propCount, len(props))
	}
}
