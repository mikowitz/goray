package goray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TransformTestCase struct {
	original, result Tuple
	transform        Matrix
	description      string
}

func TestMatrixTransformations(t *testing.T) {
	testCases := []TransformTestCase{
		TransformTestCase{
			original:    NewPoint(-3, 4, 5),
			result:      NewPoint(2, 1, 7),
			transform:   Translation(5, -3, 2),
			description: "multiplying by a translation matrix",
		},
		TransformTestCase{
			original:    NewPoint(-3, 4, 5),
			result:      NewPoint(-8, 7, 3),
			transform:   Translation(5, -3, 2).Inverse(),
			description: "multiplying by the inverse of a translation matrix",
		},
		TransformTestCase{
			original:    NewVector(-3, 4, 5),
			result:      NewVector(-3, 4, 5),
			transform:   Translation(5, -3, 2),
			description: "translation does not affect vectors",
		},
		TransformTestCase{
			transform:   Scaling(2, 3, 4),
			original:    NewPoint(-4, 6, 8),
			result:      NewPoint(-8, 18, 32),
			description: "scaling matrix applied to a point",
		},
		TransformTestCase{
			transform:   Scaling(2, 3, 4),
			original:    NewVector(-4, 6, 8),
			result:      NewVector(-8, 18, 32),
			description: "scaling matrix applied to a vector",
		},
		TransformTestCase{
			transform:   Scaling(2, 3, 4).Inverse(),
			original:    NewVector(-4, 6, 8),
			result:      NewVector(-2, 2, 2),
			description: "multiplying by the inverse of a scaling matrix",
		},
		TransformTestCase{
			transform:   Scaling(-1, 1, 1),
			original:    NewVector(2, 3, 4),
			result:      NewVector(-2, 3, 4),
			description: "reflection is scaling by a negative value",
		},
		TransformTestCase{
			transform:   RotationX(math.Pi / 4),
			original:    NewVector(0, 1, 0),
			result:      NewVector(0, math.Sqrt2/2, math.Sqrt2/2),
			description: "rotating a point π/4 around the x axis",
		},
		TransformTestCase{
			transform:   RotationX(math.Pi / 2),
			original:    NewVector(0, 1, 0),
			result:      NewVector(0, 0, 1),
			description: "rotating a point π/2 around the x axis",
		},
		TransformTestCase{
			transform:   RotationX(math.Pi / 4).Inverse(),
			original:    NewVector(0, 1, 0),
			result:      NewVector(0, math.Sqrt2/2, -math.Sqrt2/2),
			description: "the inverse of an x-rotation rotates in the opposite direction",
		},
		TransformTestCase{
			transform:   RotationY(math.Pi / 4),
			original:    NewVector(0, 0, 1),
			result:      NewVector(math.Sqrt2/2, 0, math.Sqrt2/2),
			description: "rotating a point π/4 around the y axis",
		},
		TransformTestCase{
			transform:   RotationY(math.Pi / 2),
			original:    NewVector(0, 0, 1),
			result:      NewVector(1, 0, 0),
			description: "rotating a point π/2 around the y axis",
		},
		TransformTestCase{
			transform:   RotationZ(math.Pi / 4),
			original:    NewVector(0, 1, 0),
			result:      NewVector(-math.Sqrt2/2, math.Sqrt2/2, 0),
			description: "rotating a point π/4 around the z axis",
		},
		TransformTestCase{
			transform:   RotationZ(math.Pi / 2),
			original:    NewVector(0, 1, 0),
			result:      NewVector(-1, 0, 0),
			description: "rotating a point π/2 around the z axis",
		},
		TransformTestCase{
			transform:   Shearing(1, 0, 0, 0, 0, 0),
			original:    NewVector(2, 3, 4),
			result:      NewVector(5, 3, 4),
			description: "shearing moves x in proportion to y",
		},
		TransformTestCase{
			transform:   Shearing(0, 1, 0, 0, 0, 0),
			original:    NewVector(2, 3, 4),
			result:      NewVector(6, 3, 4),
			description: "shearing moves x in proportion to z",
		},
		TransformTestCase{
			transform:   Shearing(0, 0, 1, 0, 0, 0),
			original:    NewVector(2, 3, 4),
			result:      NewVector(2, 5, 4),
			description: "shearing moves y in proportion to x",
		},
		TransformTestCase{
			transform:   Shearing(0, 0, 0, 1, 0, 0),
			original:    NewVector(2, 3, 4),
			result:      NewVector(2, 7, 4),
			description: "shearing moves y in proportion to z",
		},
		TransformTestCase{
			transform:   Shearing(0, 0, 0, 0, 1, 0),
			original:    NewVector(2, 3, 4),
			result:      NewVector(2, 3, 6),
			description: "shearing moves z in proportion to x",
		},
		TransformTestCase{
			transform:   Shearing(0, 0, 0, 0, 0, 1),
			original:    NewVector(2, 3, 4),
			result:      NewVector(2, 3, 7),
			description: "shearing moves z in proportion to y",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			fmt.Println(tc.transform)
			assert.True(t, TuplesEqual(tc.transform.Mult(tc.original), tc.result))
		})
	}
}

func TestChainingTransformations(t *testing.T) {
	t.Run("individual transformations are applied in sequence", func(t *testing.T) {
		p := NewPoint(1, 0, 1)
		a := RotationX(math.Pi / 2)
		b := Scaling(5, 5, 5)
		c := Translation(10, 5, 7)

		p2 := a.Mult(p)
		assert.True(t, TuplesEqual(p2, NewPoint(1, -1, 0)))

		p3 := b.Mult(p2)
		assert.True(t, TuplesEqual(p3, NewPoint(5, -5, 0)))

		p4 := c.Mult(p3)
		assert.True(t, TuplesEqual(p4, NewPoint(15, 0, 7)))
	})

	t.Run("chained transformations must be applied in reverse order", func(t *testing.T) {
		p := NewPoint(1, 0, 1)
		a := RotationX(math.Pi / 2)
		b := Scaling(5, 5, 5)
		c := Translation(10, 5, 7)

		transform := c.Mul(b).Mul(a)

		assert.True(t, TuplesEqual(transform.Mult(p), NewPoint(15, 0, 7)))
	})
}
