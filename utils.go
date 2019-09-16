package similarity

import "math"

// Determine float equality to 3 decimal places.
const float64EqualityThreshold = 1e-3

func almostEqual(a, b float64) bool {
    return math.Abs(a - b) <= float64EqualityThreshold
}
