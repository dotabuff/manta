package manta

import (
	"testing"
)

// Simply tests that we can read the outer structure of a real match
func TestOuterParserRealMatch(t *testing.T) {
	parser := NewParserFromFile("replays/real_match.dem")
	parser.Start()
}
