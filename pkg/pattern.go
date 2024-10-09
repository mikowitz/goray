package goray

import (
	"math"
)

type Pattern interface {
	At(p Point) Color
	AtObject(shape Shape, p Point) Color
	GetTransform() Matrix
	SetTransform(m Matrix)
}

func PatternAtObject(pattern Pattern, shape Shape, point Point) Color {
	return pattern.AtObject(shape, point)
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

func (p *SolidPattern) AtObject(shape Shape, point Point) Color {
	objectPoint := shape.GetTransform().Inverse().Mult(point)
	patternPoint := p.GetTransform().Inverse().Mult(objectPoint)
	return p.At(patternPoint)
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

func (p *StripePattern) AtObject(shape Shape, point Point) Color {
	objectPoint := shape.GetTransform().Inverse().Mult(point)
	patternPoint := p.GetTransform().Inverse().Mult(objectPoint)
	return p.At(patternPoint)
}

type GradientPattern struct {
	A, B      Color
	Transform Matrix
}

func NewGradientPattern(a, b Color) GradientPattern {
	return GradientPattern{A: a, B: b, Transform: IdentityMatrix()}
}

func (gp *GradientPattern) GetTransform() Matrix {
	return gp.Transform
}

func (gp *GradientPattern) SetTransform(m Matrix) {
	gp.Transform = m
}

func (gp *GradientPattern) At(p Point) Color {
	return gp.A.Add(gp.B.Sub(gp.A).Mul(p.x - math.Floor(p.x)))
}

func (p *GradientPattern) AtObject(shape Shape, point Point) Color {
	objectPoint := shape.GetTransform().Inverse().Mult(point)
	patternPoint := p.GetTransform().Inverse().Mult(objectPoint)
	return p.At(patternPoint)
}

type RingPattern struct {
	A, B      Color
	Transform Matrix
}

func NewRingPattern(a, b Color) RingPattern {
	return RingPattern{A: a, B: b, Transform: IdentityMatrix()}
}

func (rp *RingPattern) GetTransform() Matrix {
	return rp.Transform
}

func (rp *RingPattern) SetTransform(m Matrix) {
	rp.Transform = m
}

func (rp *RingPattern) At(p Point) Color {
	if math.Remainder(math.Sqrt(p.x*p.x+p.z*p.z), 2) == 0 {
		return rp.A
	}
	return rp.B
}

func (p *RingPattern) AtObject(shape Shape, point Point) Color {
	objectPoint := shape.GetTransform().Inverse().Mult(point)
	patternPoint := p.GetTransform().Inverse().Mult(objectPoint)
	return p.At(patternPoint)
}

type CheckersPattern struct {
	A, B      Color
	Transform Matrix
}

func NewCheckersPattern(a, b Color) CheckersPattern {
	return CheckersPattern{A: a, B: b, Transform: IdentityMatrix()}
}

func (cp *CheckersPattern) GetTransform() Matrix {
	return cp.Transform
}

func (cp *CheckersPattern) SetTransform(m Matrix) {
	cp.Transform = m
}

func (cp *CheckersPattern) At(p Point) Color {
	if math.Remainder(math.Floor(p.x)+math.Floor(p.y)+math.Floor(p.z), 2) == 0 {
		return cp.A
	}
	return cp.B
}

func (p *CheckersPattern) AtObject(shape Shape, point Point) Color {
	objectPoint := shape.GetTransform().Inverse().Mult(point)
	patternPoint := p.GetTransform().Inverse().Mult(objectPoint)
	return p.At(patternPoint)
}

type BlendedPattern struct {
	A, B      Pattern
	Transform Matrix
}

func (bp *BlendedPattern) GetTransform() Matrix {
	return bp.Transform
}

func (bp *BlendedPattern) SetTransform(m Matrix) {
	bp.Transform = m
}

func (bp *BlendedPattern) At(p Point) Color {
	return bp.A.At(p).Add(bp.B.At(p))
}

func (p *BlendedPattern) AtObject(shape Shape, point Point) Color {
	patternPoint := p.GetTransform().Inverse().Mult(point)
	return p.A.AtObject(shape, patternPoint).Add(p.B.AtObject(shape, patternPoint))
}
