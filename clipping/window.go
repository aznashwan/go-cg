package clipping

import (
	"fmt"

	"../postscript/objects"
)

// Window is the base type for a viewport object.
// The vertical minimums are modeled after the right-handed cartesian.
type Window struct {
	minx, miny, maxx, maxy int
	a, b, r, l             *objects.Line
}

// NewWindow generates a new Window instance.
func NewWindow(minx, miny, maxx, maxy int) *Window {
	return &Window{
		minx: minx,
		miny: miny,
		maxx: maxx,
		maxy: maxy,
		a:    objects.NewLine(objects.NewPoint(minx, maxy), objects.NewPoint(maxx, maxy)),
		b:    objects.NewLine(objects.NewPoint(minx, miny), objects.NewPoint(maxx, miny)),
		r:    objects.NewLine(objects.NewPoint(maxx, miny), objects.NewPoint(maxx, maxy)),
		l:    objects.NewLine(objects.NewPoint(minx, miny), objects.NewPoint(minx, maxy)),
	}
}

// ComputeABRL returns the ABRL code of a given point.
func (w *Window) ComputeABRL(p *objects.Point) int {
	var abrl int = 0

	if p.Y < w.miny {
		// A = 0; B = 1
		abrl = 1
	} else if p.Y > w.maxy {
		// A = 1; B = 0
		abrl = 2
	}

	abrl = abrl << 2

	if p.X < w.minx {
		// R = 0; L = 1
		abrl = abrl + 1
	} else if p.X > w.maxx {
		// R = 1; L = 0
		abrl = abrl + 2
	}

	return abrl
}

// LineClips returns true if the current Window clips a line.
func (w *Window) LineClips(l *objects.Line) bool {
	abrl1 := w.ComputeABRL(l.A)
	abrl2 := w.ComputeABRL(l.B)

	if abrl1|abrl2 == 0 {
		return false
	}

	if abrl1&abrl2 != 0 {
		return false
	}

	return true
}

// getLineEquation returns the parameters of the equation of the given Line.
func getLineEquation(l *objects.Line) (a, b, c int) {
	a = int(l.B.Y) - int(l.A.Y)
	b = int(l.A.X) - int(l.B.X)
	c = a*int(l.A.X) + b*int(l.A.Y)
	return
}

// computeIntersection computes the point of intersection of two given
// objects.Line objects.
func computeIntersection(L1, L2 *objects.Line) (p *objects.Point) {
	a1, b1, c1 := getLineEquation(L1)
	a2, b2, c2 := getLineEquation(L2)

	det := a1*b2 - a2*b1
	if det == 0 {
		return nil
	} else {
		x := (b2*c1 - b1*c2) / det
		y := (a1*c2 - a2*c1) / det
		p = objects.NewPoint(x, y)
	}

	return p
}

// filterIntersection is a helper method which filters which is the
// intersection of a given line with the immediate edge.
func (w *Window) filterIntersection(l *objects.Line, abrl int) *objects.Point {
	switch {
	case abrl&8 != 0:
		return computeIntersection(w.a, l)
	case abrl&4 != 0:
		return computeIntersection(w.b, l)
	case abrl&2 != 0:
		return computeIntersection(w.r, l)
	case abrl&1 != 0:
		return computeIntersection(w.l, l)
	}

	return nil
}

// ClipLine is a function which takes a Line object and returns
// a new Line which is the clipped version of the given Line.
func (w *Window) ClipLine(l *objects.Line) (*objects.Line, error) {
	abrl1 := w.ComputeABRL(l.A)
	abrl2 := w.ComputeABRL(l.B)

	if abrl1|abrl2 == 0 {
		// trivially accept
		return l, nil
	}
	if abrl1&abrl2 != 0 {
		return nil, nil
	}

	// We try the first point of the line.
	if abrl1 != 0 {
		inter := w.filterIntersection(l, abrl1)
		if inter == nil {
			return nil, fmt.Errorf("Failed to intersect line: (%d, %d) - (%d, %d)",
				l.A.X, l.A.Y, l.B.X, l.B.Y)
		}
		return w.ClipLine(objects.NewLine(inter, l.B))
	} else if abrl2 != 0 {
		inter := w.filterIntersection(l, abrl2)
		if inter == nil {
			return nil, fmt.Errorf("Failed to intersect line: (%d, %d) - (%d, %d)",
				l.A.X, l.A.Y, l.B.X, l.B.Y)
		}
		return w.ClipLine(objects.NewLine(l.A, inter))
	}

	// if no clipping is required any more:
	return l, nil
}
