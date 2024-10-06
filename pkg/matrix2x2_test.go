package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrix2x2Determinant(t *testing.T) {
	a := NewMatrix2x2(1, 5, -3, 2)

	assert.Equal(t, a.Determinant(), 17.0)
}
