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
			debug:       true,
			expectCount: 99, // tbd
		},
	}

	// Load our send tables
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables/1560315800.pbmsg"), m); err != nil {
		panic(err)
	}

	// Retrieve the flattened field serializer
	fs, err := parseSendTablesNew(m)
	assert.Nil(err)

	// Build the huffman tree
	huf := fieldpath_huffman()

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

		// Set debug status
		debugMode = s.debug

		// Initialize a field path and walk it
		fieldPath := fieldpath_init(serializer, &huf)
		fieldPath.fieldpath_walk(newReader(buf))

		// Verify field count
		assert.Equal(len(fieldPath.fields), s.expectCount)
	}
}
