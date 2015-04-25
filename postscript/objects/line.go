package objects

import (
	"fmt"

	"../../xpm"
)

// Line is the basic structure of a postscript line definition
type Line struct {
	A, B *Point
}

// String satisfies fmt.Stringer.
func (l *Line) String() string {
	return fmt.Sprintf("[%s - %s]", l.A, l.B)
}

// NewLine returns a newly generated Line structure
func NewLine(a, b *Point) *Line {
	return &Line{A: a, B: b}
}

// Draw takes an XPM image structure as parameter and proceeds to draw this
// line on it using the color code provided using Bresenham's algorithm
// The line is drawn with respect to the usual right-handed cartesian system
// If the two points of the line are outside the image, an error is returned
// NOTE: the given color code has to have been proviously added
func (l *Line) Draw(xpm *xpm.XPM, color string) error {
	// convert int coordinates to regular ints
	// for consistent operations below
	x0 := int(l.A.X)
	x1 := int(l.B.X)
	y0 := int(l.A.Y)
	y1 := int(l.B.Y)

	// compute deltax and figure out the horizontal direction
	// sx == 1 => "going" right, else left
	dx := x1 - x0
	sx := 1
	if dx < 0 {
		sx = -1
	}

	// compute deltay and figure out the vertical direction
	// sy == 1 => "going" up, else down
	dy := y1 - y0
	sy := 1
	if dy < 0 {
		sy = -1
	}

	// get absolute values of the two deltas
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	// compute err initially as the difference between our
	// two deltas. This will fluctuate as we go along
	err := dx - dy

	for {
		// draw the current pixel
		seterror := xpm.SetPixelCartesian(int(x0), int(y0), color)
		if seterror != nil {
			return seterror
		}

		// break if we're done
		if x0 == x1 && y0 == y1 {
			break
		}

		// check how we stand for our error and
		// correct accordingly
		e2 := 2 * err
		if e2 > -dy {
			err = err - dy
			x0 = x0 + sx
		}
		if e2 < dx {
			err = err + dx
			y0 = y0 + sy
		}
	}

	return nil
}
