package twod

import (
	"math"

	"../../postscript/objects"
)

// TranslatePoint applies a 2D translation with the specified parameters to the given Point.
func TranslatePoint(p *objects.Point, tx, ty int) *objects.Point {
	return getPointFromMatrix(new2DTranslationMatrix(tx, ty).Multiply(makePointMatrix(p)))
}

// TranslateLine applies a 2D translation with the specified parameters to the given Line.
func TranslateLine(l *objects.Line, tx, ty int) *objects.Line {
	return objects.NewLine(
		TranslatePoint(l.A, tx, ty),
		TranslatePoint(l.B, tx, ty),
	)
}

// new2DTranslationMatrix returns a 3x3 Matrix which represents the operation
// of 2D translation by the given parameters.
func new2DTranslationMatrix(tx, ty int) *Matrix {
	return &Matrix{
		[][]float64{
			[]float64{1, 0, float64(tx)},
			[]float64{0, 1, float64(ty)},
			[]float64{0, 0, 1},
		},
	}
}

// RotatePointAroundPoint applies a 2D rotation with the given angle to a given Point.
func RotatePointAroundPoint(p, o *objects.Point, angle int) *objects.Point {
	return getPointFromMatrix(new2DRotationMatrix(o, angle).Multiply(makePointMatrix(p)))
}

// RotateLineAroundPoint applies a 2D rotation with the given angle to a given Line.
func RotateLineAroundPoint(l *objects.Line, p *objects.Point, angle int) *objects.Line {
	return objects.NewLine(
		RotatePointAroundPoint(l.A, p, angle),
		RotatePointAroundPoint(l.B, p, angle),
	)
}

// new2DRotationMatrix returns a 3x3 Matrix which represents the operation
// of 2D rotation with the given angle around a given point.
func new2DRotationMatrix(p *objects.Point, angle int) *Matrix {
	// NOTE(aznashwan): this may be other order around...
	a := degToRadians(angle)
	rotationMatrix := &Matrix{
		[][]float64{
			[]float64{math.Cos(a), -math.Sin(a), 0},
			[]float64{math.Sin(a), math.Cos(a), 0},
			[]float64{0, 0, 1},
		},
	}
	return new2DTranslationMatrix(p.X, p.Y).Multiply(rotationMatrix).Multiply(new2DTranslationMatrix(-p.X, -p.Y))
}

// ScalePointAroundPoint applies a 2D scaling with the given factors to the given Point.
func ScalePointAroundPoint(p, o *objects.Point, sx, sy float64) *objects.Point {
	return getPointFromMatrix(new2DScalingMatrix(o, sx, sy).Multiply(makePointMatrix(p)))
}

// ScaleLineAroundPoint applies a 2D scaling with the given factors to the given Line.
func ScaleLineAroundPoint(l *objects.Line, p *objects.Point, sx, sy float64) *objects.Line {
	return objects.NewLine(
		ScalePointAroundPoint(l.A, p, sx, sy),
		ScalePointAroundPoint(l.B, p, sx, sy),
	)
}

// new2DScalingMatrix returns a 3x3 Matrix which represents the operation
// of 2D scaling with the given scaling factors around the given point.
func new2DScalingMatrix(p *objects.Point, sx, sy float64) *Matrix {
	// NOTE(aznashwa): this may be the other way around...
	scalingMatrix := &Matrix{
		[][]float64{
			[]float64{sx, 0, 0},
			[]float64{0, sy, 0},
			[]float64{0, 0, 1},
		},
	}

	return new2DTranslationMatrix(p.X, p.Y).Multiply(scalingMatrix).Multiply(new2DTranslationMatrix(-p.X, -p.Y))
}
