package goray

import (
	"fmt"
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
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			fmt.Println(tc.transform)
			assert.True(t, TuplesEqual(tc.transform.Mult(tc.original), tc.result))
		})
	}
}
