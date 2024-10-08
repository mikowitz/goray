package goray

import "math"

type Material struct {
	Ambient, Diffuse, Specular, Shininess float64
	Pattern                               Pattern
}

func NewMaterial() Material {
	pattern := NewSolidPattern(White())
	return Material{
		Pattern:   &pattern,
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (m Material) Lighting(s Shape, light PointLight, point Point, eyev, normalv Vector, inShadow bool) Color {
	effectiveColor := m.Pattern.AtObject(s, point).Prod(light.Intensity)
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

	if inShadow {
		return ambient
	}
	return ambient.Add(diffuse).Add(specular)
}
