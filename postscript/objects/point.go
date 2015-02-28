package objects

// Point is the basic datastructure represing a point by its coordinates
type Point struct {
	x, y uint
}

// NewPoint returns a newly generated Point structure
func NewPoint(x, y uint) *Point {
	return &Point{x: x, y: y}
}
