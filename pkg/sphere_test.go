package goray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSphereMaterial(t *testing.T) {
	t.Run("a sphere's default material", func(t *testing.T) {
		s := NewSphere()
		assert.Equal(t, s.Material, NewMaterial())
	})

	t.Run("changing a sphere's material", func(t *testing.T) {
		s := NewSphere()
		m := NewMaterial()
		m.Ambient = 1.0
		m.Diffuse = 0.5
		s.SetMaterial(m)
		assert.Equal(t, s.Material, m)
	})
}

func TestSphereTransform(t *testing.T) {
	t.Run("a sphere's default transformation", func(t *testing.T) {
		s := NewSphere()
		assert.True(t, MatricesEqual(s.Transform, IdentityMatrix()))
	})

	t.Run("changing a sphere's transformation", func(t *testing.T) {
		s := NewSphere()
		transform := Translation(2, 3, 4)
		s.SetTransform(transform)

		assert.True(t, MatricesEqual(s.Transform, transform))
	})
}

func TestSurfaceNormal(t *testing.T) {
	s := NewSphere()

	sqrt33 := math.Sqrt(3) / 3.0

	testCases := map[Point]Vector{
		NewPoint(1, 0, 0):                NewVector(1, 0, 0).Normalize(),
		NewPoint(0, 1, 0):                NewVector(0, 1, 0).Normalize(),
		NewPoint(0, 0, 1):                NewVector(0, 0, 1).Normalize(),
		NewPoint(sqrt33, sqrt33, sqrt33): NewVector(sqrt33, sqrt33, sqrt33).Normalize(),
	}

	for p, n := range testCases {
		assert.True(t, TuplesEqual(s.NormalAt(p), n))
	}
}

func TestNormalOnTransformedSphere(t *testing.T) {
	t.Run("on a translated sphere", func(t *testing.T) {
		s := NewSphere()
		s.SetTransform(Translation(0, 1, 0))

		n := s.NormalAt(NewPoint(0, 1.70711, -0.70711))

		assert.True(t, TuplesEqual(n, NewVector(0, 0.70711, -0.70711)))
	})

	t.Run("on a transformed sphere", func(t *testing.T) {
		s := NewSphere()
		m := Scaling(1, 0.5, 1).Mul(RotationZ(math.Pi / 5))
		s.SetTransform(m)

		n := s.NormalAt(NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2))

		assert.True(t, TuplesEqual(n, NewVector(0, 0.97014, -0.24254)))
	})
}
