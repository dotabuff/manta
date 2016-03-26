package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestFieldpath(t *testing.T) {
	assert := assert.New(t)

	// roughly the same format used in property_test.go
	scenarios := []struct {
		tableName   string // the name of the table, must have a sendtable fixture.
		run         bool   // whether or not we run the test.
		debug       bool   // whether or not we print debugging output.
		expectCount int    // how many result entries we expect.
	}{
		{
			tableName:   "CRagdollManager",
			run:         true,
			debug:       false,
			expectCount: 1,
		},
		{
			tableName:   "CDOTATeam",
			run:         true,
			debug:       false,
			expectCount: 15,
		},
		{
			tableName:   "CWorld",
			run:         true,
			debug:       false,
			expectCount: 139,
		},
		{
			tableName:   "CDOTAPlayer",
			run:         true,
			debug:       false,
			expectCount: 137,
		},
		{
			tableName:   "CDOTA_PlayerResource",
			run:         true,
			debug:       false,
			expectCount: 2056,
		},
		{
			tableName:   "CBaseAnimating",
			run:         true,
			debug:       false,
			expectCount: 110,
		},
		{
			tableName:   "CBaseEntity",
			run:         true,
			debug:       false,
			expectCount: 35,
		},
		{
			tableName:   "CDOTAGamerulesProxy",
			run:         true,
			debug:       false,
			expectCount: 1838,
		},
		{
			tableName:   "CSpeechBubbleManager",
			run:         true,
			debug:       false,
			expectCount: 25,
		},
	}

	// Load our send tables
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables/1560315800.pbmsg"), m); err != nil {
		panic(err)
	}

	// Retrieve the flattened field serializer
	p := &Parser{}
	fs := p.parseSendTables(m, newPropertySerializerTable())

	// Iterate over the different scenarios
	// -! Create a new FieldPath for each scenario
	for _, s := range scenarios {
		// Load up a fixture
		buf := _read_fixture(_sprintf("instancebaseline/1560315800_%s.rawbuf", s.tableName))

		// Get the serializer
		// We don't really know which version is used to generate the baseline
		// 0 seems resonable
		serializer := fs.Serializers[s.tableName][0]
		assert.NotNil(serializer)

		// Optionally skip
		if !s.run {
			continue
		}

		// Initialize a field path and walk it
		fieldPath := newFieldpath(serializer)
		fieldPath.walk(newReader(buf))

		// Verify field count
		assert.Equal(len(fieldPath.fields), s.expectCount)

		// Print a list of all fields read
		for i, f := range fieldPath.fields {
			if f.Field.Index >= 0 {
				_debugf("%d\t%s[%d]\t%s", i, f.Name, f.Field.Index, f.Field.Type)
			} else {
				_debugf("%d\t%s\t%s", i, f.Name, f.Field.Type)
			}
		}
	}
}
