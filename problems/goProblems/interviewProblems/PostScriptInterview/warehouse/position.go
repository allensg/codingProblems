package warehouse

import (
	"fmt"
	"math"
)

// Position represents an x, y coordinate in a given warehouse. Position can be used to determine how far apart or near together two vehicles are.
type Position struct {
	X, Y float64
}

// String returns the string representation of a Position.
func (p *Position) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}

// Equal returns true if the input Position represents the same location.
func (p *Position) Equal(q *Position) bool {
	return p.X == q.X && p.Y == q.Y
}

// Distance returns the distance between two Positions. It is calculated as the Euclidean distance in two dimensions (https://en.wikipedia.org/wiki/Euclidean_distance).
func Distance(p, q *Position) float64 {
	xDiff := math.Abs(p.X - q.X)
	yDiff := math.Abs(p.Y - q.Y)
	return math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
}
