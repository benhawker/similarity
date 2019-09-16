package similarity

import (
	"math"
)

// Manhattan distance is a metric in which the distance between two points is 
// the sum of the absolute differences of their Cartesian coordinates.
// In a simple way of saying it is the total suzm of the difference between 
// the x-coordinates  and y-coordinates.

// Suppose we have two points A and B if we want to find the Manhattan distance 
// between them, just we have, to sum up, the absolute x-axis and y – axis variation 
// means we have to find how these two points A and B are varying in X-axis and Y- axis.
// In a more mathematical way of saying Manhattan distance between two points measured 
// along axes at right angles.

// In a plane with p1 at (x1, y1) and p2 at (x2, y2).
// Manhattan distance = |x1 – x2| + |y1 – y2|

// This Manhattan distance metric is also known as
// 	- Manhattan length
// 	- rectilinear distance
// 	- L1 distance or L1 norm
// 	- city block distance
// 	- Minkowski’s L1 distance
// 	- taxi-cab metric
// 	- or city block distance.
func Manhattan(coords1, coords2 []float64) (float64, error) {
	sum := 0.0

	for index, coord := range coords1 {

		sum += math.Abs(coords2[index] - coord)
	}

	return sum, nil
}
