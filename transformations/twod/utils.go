package twod

import (
	"math"

	"../../postscript/objects"
)

// degToRadians converts the given number of degrees to radians.
func degToRadians(degs int) float64 {
	return float64(degs) * 180 / math.Pi
}

// makePointMatrix takes a point and makes its respective column matrix.
func makePointMatrix(p *objects.Point) *Matrix {
	return &Matrix{
		[][]float64{
			[]float64{float64(p.X)},
			[]float64{float64(p.Y)},
			[]float64{1},
		},
	}
}

// getPointFromMatrix extracts the value of a Point from its characteristic matrix.
func getPointFromMatrix(m *Matrix) *objects.Point {
	return objects.NewPoint(int(m.rows[0][0]), int(m.rows[1][0]))
}
