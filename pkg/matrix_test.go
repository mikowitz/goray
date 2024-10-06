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

func TestMultiplyingMatrices(t *testing.T) {
	t.Run("multiplying by a matrix", func(t *testing.T) {
		a := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
		b := NewMatrix(-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8)

		expected := NewMatrix(
			20, 22, 50, 48,
			44, 54, 114, 108,
			40, 58, 110, 102,
			16, 26, 46, 42,
		)

		assert.True(t, MatricesEqual(a.Mul(b), expected))
	})

	t.Run("multiplying by a tuple", func(t *testing.T) {
		a := NewMatrix(1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1)
		tup := NewTuple(1, 2, 3, 1)

		assert.True(t, TuplesEqual(a.Mult(tup), NewTuple(18, 24, 33, 1)))
	})

	t.Run("multiplying by the identity matrix", func(t *testing.T) {
		a := NewMatrix(0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32)

		assert.True(t, MatricesEqual(a.Mul(IdentityMatrix()), a))
	})

	t.Run("multiplying the identity matrix by a tuple", func(t *testing.T) {
		a := NewTuple(1, 2, 3, 4)

		assert.True(t, TuplesEqual(IdentityMatrix().Mult(a), a))
	})
}

func TestTransposingMatrices(t *testing.T) {
	t.Run("transposing a matrix", func(t *testing.T) {
		a := NewMatrix(0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8)
		expected := NewMatrix(0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8)

		assert.True(t, MatricesEqual(a.Transpose(), expected))
	})
}

func TestSubmatrixOfMatrix(t *testing.T) {
	a := NewMatrix(-6, 1, 1, 6, -8, 5, 8, 6, -1, 0, 8, 2, -7, 1, -1, 1)
	expected := NewMatrix3x3(-6, 1, 6, -8, 8, 6, -7, -1, 1)

	assert.True(t, Matrices33Equal(a.Submatrix(2, 1), expected))
}

func TestDeterminantOfMatrix(t *testing.T) {
	a := NewMatrix(-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9)

	assert.Equal(t, a.Cofactor(0, 0), 690.0)
	assert.Equal(t, a.Cofactor(0, 1), 447.0)
	assert.Equal(t, a.Cofactor(0, 2), 210.0)
	assert.Equal(t, a.Cofactor(0, 3), 51.0)
	assert.Equal(t, a.Determinant(), -4071.0)
}

func TestMatrixInvertibility(t *testing.T) {
	t.Run("for an invertible matrix", func(t *testing.T) {
		a := NewMatrix(6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6)

		assert.Equal(t, a.Determinant(), -2120.0)
		assert.True(t, a.IsInvertible())
	})

	t.Run("for an invertible matrix", func(t *testing.T) {
		a := NewMatrix(-4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0)

		assert.Equal(t, a.Determinant(), 0.0)
		assert.False(t, a.IsInvertible())
	})
}

func TestMatrixInverse(t *testing.T) {
	t.Run("calculating the inverse of a matrix", func(t *testing.T) {
		a := NewMatrix(-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4)
		b := a.Inverse()

		assert.Equal(t, a.Determinant(), 532.0)
		assert.Equal(t, a.Cofactor(2, 3), -160.0)
		assert.Equal(t, b.At(3, 2), -160.0/532.0)
		assert.Equal(t, a.Cofactor(3, 2), 105.0)
		assert.Equal(t, b.At(2, 3), 105.0/532.0)

		expected := NewMatrix(0.21805, 0.45113, 0.24060, -0.04511, -0.80827, -1.45677, -0.44361, 0.52068, -0.07895, -0.22368, -0.05263, 0.19737, -0.52256, -0.81391, -0.30075, 0.30639)

		assert.True(t, MatricesEqual(b, expected))

	})

	t.Run("calculating the inverse of another matrix", func(t *testing.T) {
		a := NewMatrix(8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4)

		expected := NewMatrix(-0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308, 0.02564, 0.03077, 0.35897, 0.35897, 0.43590, 0.92308, -0.69231, -0.69231, -0.76923, -1.92308)

		assert.True(t, MatricesEqual(a.Inverse(), expected))
	})

	t.Run("calculating the inverse of a third matrix", func(t *testing.T) {
		a := NewMatrix(9, 3, 0, 9, -5, -2, -6, -3, -4, 9, 6, 4, -7, 6, 6, 2)

		expected := NewMatrix(-0.04074, -0.07778, 0.14444, -0.22222, -0.07778, 0.03333, 0.36667, -0.33333, -0.02901, -0.14630, -0.10926, 0.12963, 0.17778, 0.06667, -0.26667, 0.33333)

		assert.True(t, MatricesEqual(a.Inverse(), expected))
	})

	t.Run("multiplying a matrix by its inverse", func(t *testing.T) {
		a := NewMatrix(3, -9, 7, 3, 3, -8, 2, -9, -4, 4, 4, 1, -6, 5, -1, 1)
		b := NewMatrix(8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5, 4, 6, -2, 0, 5)

		c := a.Mul(b)

		assert.True(t, MatricesEqual(c.Mul(b.Inverse()), a))
	})
}
