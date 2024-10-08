package goray

import "math"

type Pattern interface {
	At(p Point) Color
	AtObject(s Shape, p Point) Color
	GetTransform() Matrix
	SetTransform(m Matrix)
}

type SolidPattern struct {
	Color     Color
	Transform Matrix
}

func NewSolidPattern(c Color) SolidPattern {
	return SolidPattern{Color: c, Transform: IdentityMatrix()}
}

func (sp *SolidPattern) GetTransform() Matrix {
	return sp.Transform
}

func (sp *SolidPattern) SetTransform(m Matrix) {
	sp.Transform = m
}

func (sp *SolidPattern) At(_ Point) Color {
	return sp.Color
}

func (sp *SolidPattern) AtObject(_ Shape, _ Point) Color {
	return sp.Color
}

type StripePattern struct {
	A, B      Color
	Transform Matrix
}

func NewStripePattern(a, b Color) StripePattern {
	return StripePattern{A: a, B: b, Transform: IdentityMatrix()}
}

func (sp *StripePattern) GetTransform() Matrix {
	return sp.Transform
}

func (sp *StripePattern) SetTransform(m Matrix) {
	sp.Transform = m
}

func (sp *StripePattern) At(p Point) Color {
	if math.Remainder(math.Floor(p.x), 2) == 0 {
		return sp.A
	}
	return sp.B
}

func (sp *StripePattern) AtObject(s Shape, p Point) Color {
	objectPoint := s.GetTransform().Inverse().Mult(p)
	patternPoint := sp.Transform.Inverse().Mult(objectPoint)
	return sp.At(patternPoint)
}
