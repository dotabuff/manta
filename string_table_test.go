package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestStringTableParsing(t *testing.T) {
	assert := assert.New(t)

	// A message with no data
	m := &dota.CSVCMsg_CreateStringTable{}
	if err := proto.Unmarshal(_read_fixture("string_tables_1_29_uncompressed.pbmsg"), m); err != nil {
		panic(err)
	}
	// Verify table name
	assert.Equal("downloadables", m.GetName())
	assert.Equal(0, len(m.GetStringData()))
	// Verify no items are parsed
	items := parseStringTable(m)
	assert.Equal(0, len(items))

	// A compressed message with 4162 bytes of data
	m = &dota.CSVCMsg_CreateStringTable{}
	if err := proto.Unmarshal(_read_fixture("string_tables_8_4162_compressed.pbmsg"), m); err != nil {
		panic(err)
	}
	assert.Equal("EntityNames", m.GetName())
	assert.Equal(4162, len(m.GetStringData()))
	assert.Equal(true, m.GetDataCompressed())

	// An uncompressed message with 559 bytes of data
	m = &dota.CSVCMsg_CreateStringTable{}
	if err := proto.Unmarshal(_read_fixture("string_tables_16_559_uncompressed.pbmsg"), m); err != nil {
		panic(err)
	}
	assert.Equal("EconItems", m.GetName())
	assert.Equal(533, len(m.GetStringData()))

	// XXX TODO: successfully parse one of these things!
}
