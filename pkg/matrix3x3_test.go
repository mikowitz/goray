package goray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Matrices22Equal(m, n Matrix2x2) bool {
	ε := 0.00001
	for i, _ := range m {
		if math.Abs(m[i]-n[i]) > ε {
			return false
		}
	}
	return true
}

func Matrices33Equal(m, n Matrix3x3) bool {
	ε := 0.00001
	for i, _ := range m {
		if math.Abs(m[i]-n[i]) > ε {
			return false
		}
	}
	return true
}

func TestSubmatrixOf3x3Matrix(t *testing.T) {
	a := NewMatrix3x3(1, 5, 0, -3, 2, 7, 0, 6, -3)
	expected := NewMatrix2x2(-3, 2, 0, 6)

	assert.True(t, Matrices22Equal(a.Submatrix(0, 2), expected))
}

func TestMinorOf3x3Matrix(t *testing.T) {
	a := NewMatrix3x3(3, 5, 0, 2, -1, -7, 6, -1, 5)
	b := a.Submatrix(1, 0)

	assert.Equal(t, b.Determinant(), 25.0)
	assert.Equal(t, a.Minor(1, 0), 25.0)
}

func TestCofactorOf3x3Matrix(t *testing.T) {
	a := NewMatrix3x3(3, 5, 0, 2, -1, -7, 6, -1, 5)

	assert.Equal(t, a.Minor(0, 0), -12.0)
	assert.Equal(t, a.Cofactor(0, 0), -12.0)

	assert.Equal(t, a.Minor(1, 0), 25.0)
	assert.Equal(t, a.Cofactor(1, 0), -25.0)
}

func TestDeterminantOf3x3Matrix(t *testing.T) {
	a := NewMatrix3x3(1, 2, 6, -5, 8, -4, 2, 6, 4)

	assert.Equal(t, a.Cofactor(0, 0), 56.0)
	assert.Equal(t, a.Cofactor(0, 1), 12.0)
	assert.Equal(t, a.Cofactor(0, 2), -46.0)
	assert.Equal(t, a.Determinant(), -196.0)
}
