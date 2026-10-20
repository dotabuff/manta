package manta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCloneIsolatesMutableLeaves verifies entities cloned from a baseline
// template cannot affect the template or each other: not by mutating a vector
// or blob returned from get (shared leaves are handed out as copies), and not
// by setting values (including in nested states).
func TestCloneIsolatesMutableLeaves(t *testing.T) {
	assert := assert.New(t)

	at := func(path ...int32) *fieldPath {
		fp := &fieldPath{}
		copy(fp.path[:], path)
		fp.last = int32(len(path) - 1)
		return fp
	}

	tmpl := newFieldState()
	tmpl.set(at(0), floatsCell([]float32{1, 2, 3}))
	tmpl.set(at(1), bytesCell([]byte{4, 5}))
	tmpl.set(at(2, 0), floatsCell([]float32{6, 7}))
	tmpl.set(at(3), int32Cell(8))

	a, b := tmpl.clone(), tmpl.clone()

	// Mutate everything handed out by a.
	a.get(at(0)).([]float32)[0] = 100
	a.get(at(1)).([]byte)[0] = 100
	a.get(at(2, 0)).([]float32)[0] = 100
	a.set(at(3), int32Cell(100))
	a.set(at(2, 1), int32Cell(100))

	for name, s := range map[string]*fieldState{"template": tmpl, "sibling": b} {
		assert.Equal([]float32{1, 2, 3}, s.get(at(0)), name)
		assert.Equal([]byte{4, 5}, s.get(at(1)), name)
		assert.Equal([]float32{6, 7}, s.get(at(2, 0)), name)
		assert.Equal(int32(8), s.get(at(3)), name)
		assert.Nil(s.get(at(2, 1)), name)
	}

	// The clone's own shared leaves are unchanged too; its sets took effect.
	assert.Equal([]float32{1, 2, 3}, a.get(at(0)))
	assert.Equal([]byte{4, 5}, a.get(at(1)))
	assert.Equal([]float32{6, 7}, a.get(at(2, 0)))
	assert.Equal(int32(100), a.get(at(3)))
	assert.Equal(int32(100), a.get(at(2, 1)))

	// A freshly decoded value replaces the shared mark.
	v := []float32{9, 9}
	a.set(at(0), floatsCell(v))
	v[0] = 10
	assert.Equal([]float32{10, 9}, a.get(at(0)))
}
