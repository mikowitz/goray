package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayAt(t *testing.T) {
	r := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))

	assert.True(t, TuplesEqual(r.At(0), NewPoint(2, 3, 4)))
	assert.True(t, TuplesEqual(r.At(1), NewPoint(3, 3, 4)))
	assert.True(t, TuplesEqual(r.At(-1), NewPoint(1, 3, 4)))
	assert.True(t, TuplesEqual(r.At(2.5), NewPoint(4.5, 3, 4)))
}
