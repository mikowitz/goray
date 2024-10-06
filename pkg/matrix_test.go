package goray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MatricesEqual(m, n Matrix) bool {
	ε := 0.00001
	for i, _ := range m {
		if math.Abs(m[i]-n[i]) > ε {
			return false
		}
	}
	return true
}

func TestConstructingA4x4Matrix(t *testing.T) {
	m := NewMatrix(
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	)

	assert.Equal(t, len(m), 16)
	assert.Equal(t, m.At(0, 0), 1.0)
	assert.Equal(t, m.At(0, 3), 4.0)
	assert.Equal(t, m.At(1, 0), 5.5)
	assert.Equal(t, m.At(1, 2), 7.5)
	assert.Equal(t, m.At(2, 2), 11.0)
	assert.Equal(t, m.At(3, 0), 13.5)
	assert.Equal(t, m.At(3, 2), 15.5)
}

func TestConstructingSmallerMatrices(t *testing.T) {
	t.Run("a 2x2 matrix", func(t *testing.T) {
		m := NewMatrix2x2(-3, 5, 1, -2)

		assert.Equal(t, m.At(0, 0), -3.0)
		assert.Equal(t, m.At(0, 1), 5.0)
		assert.Equal(t, m.At(1, 0), 1.0)
		assert.Equal(t, m.At(1, 1), -2.0)
	})

	t.Run("a 3x3 matrix", func(t *testing.T) {
		m := NewMatrix3x3(-3, 5, 0, 1, -2, -7, 0, 1, 1)

		assert.Equal(t, m.At(0, 0), -3.0)
		assert.Equal(t, m.At(1, 1), -2.0)
		assert.Equal(t, m.At(2, 2), 1.0)
	})
}

func TestMatrixEquality(t *testing.T) {
	a := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	b := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	c := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1, 0)

	assert.True(t, MatricesEqual(a, b))
	assert.False(t, MatricesEqual(a, c))
}
