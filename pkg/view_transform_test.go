package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ViewTransformTestCase struct {
	from, to    Point
	up          Vector
	expected    Matrix
	description string
}

func TestViewTransform(t *testing.T) {
	testCases := []ViewTransformTestCase{
		ViewTransformTestCase{
			from:        NewPoint(0, 0, 0),
			to:          NewPoint(0, 0, -1),
			up:          NewVector(0, 1, 0),
			expected:    IdentityMatrix(),
			description: "for the default orientation",
		},
		ViewTransformTestCase{
			from:        NewPoint(0, 0, 0),
			to:          NewPoint(0, 0, 1),
			up:          NewVector(0, 1, 0),
			expected:    Scaling(-1, 1, -1),
			description: "looking in the positive z direction",
		},
		ViewTransformTestCase{
			from:        NewPoint(0, 0, 8),
			to:          NewPoint(0, 0, 0),
			up:          NewVector(0, 1, 0),
			expected:    Translation(0, 0, -8),
			description: "moves the world",
		},
		ViewTransformTestCase{
			from:        NewPoint(1, 3, 2),
			to:          NewPoint(4, -2, 8),
			up:          NewVector(1, 1, 0),
			expected:    NewMatrix(-0.50709, 0.50709, 0.67612, -2.36643, 0.76772, 0.60609, 0.12122, -2.82843, -0.35857, 0.59761, -0.71714, 0.00000, 0.00000, 0.00000, 0.00000, 1.00000),
			description: "arbitrary",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			vt := NewViewTransform(tc.from, tc.to, tc.up)
			assert.True(t, MatricesEqual(vt, tc.expected))
		})
	}

}
