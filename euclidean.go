package similarity

import (
	"math"
)

// In mathematics, the Euclidean distance or Euclidean metric is the "ordinary"
// straight-line distance between two points in Euclidean space.
// With this distance, Euclidean space becomes a metric space.
// The associated norm is called the Euclidean norm.
// A generalized term for the Euclidean norm is the L2 norm or L2 distance.
func Euclidean(coords1, coords2 []float64) (float64, error) {
	var sumOfSquares float64

	for index, coord := range coords1 {
		sumOfSquares += math.Pow((coord - coords2[index]), 2.0)
	}

	return math.Sqrt(sumOfSquares), nil
}
