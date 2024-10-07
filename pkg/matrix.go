package goray

type Matrix [16]float64

func NewMatrix(m ...float64) Matrix {
	return Matrix(m[:16])
}

func IdentityMatrix() Matrix {
	return Matrix{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

func (m *Matrix) Write(row, col int, value float64) {
	m[row*4+col] = value
}

func (m Matrix) At(row, col int) float64 {
	return m[row*4+col]
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

func (m Matrix) Transpose() Matrix {
	r := make([]float64, 16)
	for row := range 4 {
		for col := range 4 {
			r[col*4+row] = m.At(row, col)
		}
	}
	return Matrix(r)
}

func (m Matrix) Submatrix(xrow, xcol int) Matrix3x3 {
	r := make([]float64, 0)
	for row := range 4 {
		for col := range 4 {
			if row != xrow && col != xcol {
				r = append(r, m.At(row, col))
			}
		}
	}
	return Matrix3x3(r)
}

func (m Matrix) Minor(row, col int) float64 {
	return m.Submatrix(row, col).Determinant()
}

func (m Matrix) Cofactor(row, col int) float64 {
	minor := m.Minor(row, col)
	if (row+col)%2 == 1 {
		return -minor
	}
	return minor
}

func (m Matrix) Determinant() float64 {
	r := 0.0
	for col := range 4 {
		r += m.At(0, col) * m.Cofactor(0, col)
	}
	return r
}

func (m Matrix) IsInvertible() bool {
	return m.Determinant() != 0.0
}

func (m Matrix) Inverse() Matrix {
	if !m.IsInvertible() {
		panic(m)
	}
	r := make([]float64, 16)
	for row := range 4 {
		for col := range 4 {
			r[col*4+row] = m.Cofactor(row, col) / m.Determinant()
		}
	}
	return Matrix(r)
}
