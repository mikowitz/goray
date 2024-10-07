package goray

type PointLight struct {
	Position  Point
	Intensity Color
}

func NewPointLight(position Point, intensity Color) PointLight {
	return PointLight{Position: position, Intensity: intensity}
}
