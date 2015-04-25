package objects

import "fmt"

// Point is the basic datastructure represing a point by its coordinates
type Point struct {
	X, Y int
}

// String satisfies fmt.Stringer.
func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// NewPoint returns a newly generated Point structure
func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}
