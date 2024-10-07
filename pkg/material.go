package goray

import "math"

type Material struct {
	Color                                 Color
	Ambient, Diffuse, Specular, Shininess float64
}

func NewMaterial() Material {
	return Material{
		Color:     NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (m Material) Lighting(light PointLight, point Point, eyev, normalv Vector) Color {
	effectiveColor := m.Color.Prod(light.Intensity)
	lightv := light.Position.Sub(point).Normalize()

	ambient := effectiveColor.Mul(m.Ambient)
	diffuse := NewColor(0, 0, 0)
	specular := NewColor(0, 0, 0)

	lightDotNormal := lightv.Dot(normalv)

	if lightDotNormal >= 0 {
		diffuse = effectiveColor.Mul(m.Diffuse * lightDotNormal)

		reflectv := lightv.Neg().Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)

		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Mul(m.Specular * factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
