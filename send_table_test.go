package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestSendTableParsing(t *testing.T) {
	assert := assert.New(t)

	// The single message from a real match
	m := &dota.CDemoSendTables{}
	if err := proto.Unmarshal(_read_fixture("send_tables_1.pbmsg"), m); err != nil {
		panic(err)
	}

	// Just a basic length check
	assert.Equal(185130, len(m.GetData()))

	// XXX TODO: we'll want to do something with this.
	err := parseSendTables(m)
	assert.Nil(err)
}
