package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointLight(t *testing.T) {
	intensity := NewColor(1, 1, 1)
	position := NewPoint(0, 0, 0)

	light := NewPointLight(position, intensity)

	assert.True(t, TuplesEqual(light.Position, position))
	assert.True(t, TuplesEqual(light.Intensity, intensity))
}
