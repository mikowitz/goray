package goray

type Matrix [16]float64
type Matrix3x3 [9]float64
type Matrix2x2 [4]float64

func NewMatrix(m ...float64) Matrix {
	return Matrix(m[:16])
}

func NewMatrix3x3(m ...float64) Matrix3x3 {
	return Matrix3x3(m[:9])
}

func NewMatrix2x2(m ...float64) Matrix2x2 {
	return Matrix2x2(m[:4])
}

func (m Matrix) At(row, col int) float64 {
	return m[row*4+col]
}

func (m Matrix3x3) At(row, col int) float64 {
	return m[row*3+col]
}

func (m Matrix2x2) At(row, col int) float64 {
	return m[row*2+col]
}

func (m Matrix) Mul(n Matrix) Matrix {
	r := make([]float64, 16)

	for row := range 4 {
		for col := range 4 {
			val := 0.0
			for i := range 4 {
				val += m.At(row, i) * n.At(i, col)
			}
			r[row*4+col] = val
		}
	}
	return Matrix(r)
}

func (m Matrix) Mult(t Tuple) Tuple {
	r := make([]float64, 4)
	for i := range 4 {
		s := i * 4
		t2 := NewTuple(m[s], m[s+1], m[s+2], m[s+3])
		r[i] = t.Dot(t2)
	}
	return NewTuple(r[0], r[1], r[2], r[3])
}
