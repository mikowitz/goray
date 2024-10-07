package goray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMaterial(t *testing.T) {
	m := NewMaterial()

	assert.Equal(t, m.Color, NewColor(1, 1, 1))
	assert.Equal(t, m.Ambient, 0.1)
	assert.Equal(t, m.Diffuse, 0.9)
	assert.Equal(t, m.Specular, 0.9)
	assert.Equal(t, m.Shininess, 200.0)
}

type LightingTestCase struct {
	description   string
	eyev, normalv Vector
	light         PointLight
	result        Color
}

func TestLighting(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)

	testCases := []LightingTestCase{
		LightingTestCase{
			description: "with the eye between the light and surface",
			eyev:        NewVector(0, 0, -1),
			normalv:     NewVector(0, 0, -1),
			light:       NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1)),
			result:      NewColor(1.9, 1.9, 1.9),
		},
		LightingTestCase{
			description: "with the eye between the light and surface. eye offset 45 degrees",
			eyev:        NewVector(0, math.Sqrt2/2, -math.Sqrt2/2),
			normalv:     NewVector(0, 0, -1),
			light:       NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1)),
			result:      NewColor(1.0, 1.0, 1.0),
		},
		LightingTestCase{
			description: "with the eye opposite surface, light offset 45 degrees",
			eyev:        NewVector(0, 0, -1),
			normalv:     NewVector(0, 0, -1),
			light:       NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1)),
			result:      NewColor(0.7364, 0.7364, 0.7364),
		},
		LightingTestCase{
			description: "with the eye in the path of the reflection vector",
			eyev:        NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2),
			normalv:     NewVector(0, 0, -1),
			light:       NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1)),
			result:      NewColor(1.6364, 1.6364, 1.6364),
		},
		LightingTestCase{
			description: "with the light behind the surface",
			eyev:        NewVector(0, 0, -1),
			normalv:     NewVector(0, 0, -1),
			light:       NewPointLight(NewPoint(0, 10, 10), NewColor(1, 1, 1)),
			result:      NewColor(0.1, 0.1, 0.1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := m.Lighting(tc.light, position, tc.eyev, tc.normalv)
			assert.True(t, TuplesEqual(tc.result, actual))
		})
	}
}
