package goray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TuplesEqual(a, b Tuple) bool {
	ε := 0.00001
	return math.Abs(a.x-b.x) < ε &&
		math.Abs(a.y-b.y) < ε &&
		math.Abs(a.z-b.z) < ε &&
		math.Abs(a.w-b.w) < ε
}

func TestTuplePredicates(t *testing.T) {
	t.Run("a tuple with w=1 is a point", func(t *testing.T) {
		a := NewTuple(4.3, -4.2, 3.1, 1)

		assert.True(t, a.IsPoint())
		assert.False(t, a.IsVector())
	})

	t.Run("a tuple with w=0 is a vector", func(t *testing.T) {
		a := NewTuple(4.3, -4.2, 3.1, 0)

		assert.True(t, a.IsVector())
		assert.False(t, a.IsPoint())
	})
}

func TestNewPoint(t *testing.T) {
	p := NewPoint(4, -4, 3)
	assert.True(t, p.IsPoint())
}

func TestNewVector(t *testing.T) {
	p := NewVector(4, -4, 3)
	assert.True(t, p.IsVector())
}

func TestAddingTuples(t *testing.T) {
	a := NewTuple(3, -2, 5, 1)
	b := NewTuple(-2, 3, 1, 0)

	assert.True(t, TuplesEqual(a.Add(b), NewTuple(1, 1, 6, 1)))
}

func TestSubtractingTuples(t *testing.T) {
	t.Run("subtracting two points", func(t *testing.T) {
		p1 := NewPoint(3, 2, 1)
		p2 := NewPoint(5, 6, 7)

		assert.True(t, TuplesEqual(p1.Sub(p2), NewVector(-2, -4, -6)))
	})

	t.Run("subtracting a vector from a point", func(t *testing.T) {
		p := NewPoint(3, 2, 1)
		v := NewVector(5, 6, 7)

		assert.True(t, TuplesEqual(p.Sub(v), NewPoint(-2, -4, -6)))
	})

	t.Run("subtracting two vectors", func(t *testing.T) {
		v1 := NewVector(3, 2, 1)
		v2 := NewVector(5, 6, 7)

		assert.True(t, TuplesEqual(v1.Sub(v2), NewVector(-2, -4, -6)))
	})
}

func TestNegatingTuples(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)

	assert.True(t, TuplesEqual(a.Neg(), NewTuple(-1, 2, -3, 4)))
}

func TestMultiplyingTuples(t *testing.T) {
	t.Run("multiplying a tuple by a scalar", func(t *testing.T) {
		a := NewTuple(1, -2, 3, -4)
		assert.True(t, TuplesEqual(a.Mul(3.5), NewTuple(3.5, -7, 10.5, -14)))
	})

	t.Run("multiplying a tuple by a fraction", func(t *testing.T) {
		a := NewTuple(1, -2, 3, -4)
		assert.True(t, TuplesEqual(a.Mul(.5), NewTuple(.5, -1, 1.5, -2)))
	})
}

func TestDividingTuples(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)
	assert.True(t, TuplesEqual(a.Div(2), NewTuple(.5, -1, 1.5, -2)))
}

func TestMagnitude(t *testing.T) {
	testCases := map[Tuple]float64{
		NewVector(1, 0, 0):    1,
		NewVector(0, 1, 0):    1,
		NewVector(0, 0, 1):    1,
		NewVector(1, 2, 3):    math.Sqrt(14),
		NewVector(-1, -2, -3): math.Sqrt(14),
	}

	for v, m := range testCases {
		assert.Equal(t, v.Magnitude(), m)
	}
}

func TestNormalize(t *testing.T) {
	testCases := map[Vector]Vector{
		NewVector(4, 0, 0): NewVector(1, 0, 0),
		NewVector(1, 2, 3): NewVector(0.26726, 0.53452, 0.80178),
	}

	for v, n := range testCases {
		assert.True(t, TuplesEqual(v.Normalize(), n))
	}
}

func TestDotProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	assert.Equal(t, a.Dot(b), 20.0)
}

func TestCrossProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	assert.True(t, TuplesEqual(a.Cross(b), NewVector(-1, 2, -1)))
	assert.True(t, TuplesEqual(b.Cross(a), NewVector(1, -2, 1)))
}

func TestReflect(t *testing.T) {
	t.Run("reflecting a vector approaching at 45 degrees", func(t *testing.T) {
		v := NewVector(1, -1, 0)
		n := NewVector(0, 1, 0)

		r := v.Reflect(n)

		assert.True(t, TuplesEqual(r, NewVector(1, 1, 0)))
	})

	t.Run("reflecting a vector off a slanted surface", func(t *testing.T) {
		v := NewVector(0, -1, 0)
		n := NewVector(math.Sqrt2/2, math.Sqrt2/2, 0)

		r := v.Reflect(n)

		assert.True(t, TuplesEqual(r, NewVector(1, 0, 0)))
	})
}
