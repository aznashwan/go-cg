package twod

import (
	"fmt"
)

// Matrix is the basic 2d transformation matrix object.
type Matrix struct {
	rows [][]float64
}

// NewMatrix returns a new fully initialized Matrix.
func NewMatrix(nrows, ncols int) *Matrix {
	rows := make([][]float64, nrows)
	for i := 0; i < nrows; i++ {
		rows[i] = make([]float64, ncols)
	}
	return &Matrix{rows}
}

// String satisfies the io.Stringer interface.
func (m *Matrix) String() string {
	res := ""
	for _, row := range m.rows {
		res = res + "| "
		for _, n := range row {
			res = res + fmt.Sprintf("%d ", int(n))
		}
		res = res[:len(res)] + "|\n"
	}
	return res
}

// Multiply multiplies this Matrix struct with another:
func (m *Matrix) Multiply(other *Matrix) *Matrix {
	m1 := len(m.rows)
	n1 := len(m.rows[0])
	m2 := len(other.rows)
	n2 := len(other.rows[0])

	if n1 != m2 {
		panic(fmt.Sprintf("Matrices are incompatible for multimplication:\n%s\n%s", m.String(), other.String()))
	}

	res := NewMatrix(m1, n2)

	for i := 0; i < m1; i++ {
		for j := 0; j < n2; j++ {
			res.rows[i][j] = 0
			for k := 0; k < n1; k++ {
				res.rows[i][j] = res.rows[i][j] + m.rows[i][k]*other.rows[k][j]
			}
		}
	}

	return res
}
