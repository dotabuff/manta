package manta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityOpFlag(t *testing.T) {
	assert := assert.New(t)

	assert.True(EntityOpCreated.Flag(EntityOpCreated))
	assert.False(EntityOpCreated.Flag(EntityOpDeleted))
	assert.False(EntityOpCreated.Flag(EntityOpEntered))
	assert.True(EntityOpCreatedEntered.Flag(EntityOpCreated))
	assert.True(EntityOpCreatedEntered.Flag(EntityOpEntered))
	assert.False(EntityOpCreatedEntered.Flag(EntityOpDeleted))
	assert.False(EntityOpCreatedEntered.Flag(EntityOpLeft))
}
