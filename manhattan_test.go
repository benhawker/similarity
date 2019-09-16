package similarity

import (
	"testing"
)

func TestManhattan(t *testing.T) {
	// Same coords
	a := []float64{1, 1}
	b := []float64{1, 1}
	
	distance, err := Manhattan(a, b)
	if err != nil {
		t.Error(err)
	}
	if distance != 0 {
		t.Error("Expected distance of 0, got instead ", distance)
	}

	// Different coords with 3 dimensions
	a = []float64{1, 2, 3}
	b = []float64{4, 5, 6}

	distance, err = Manhattan(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 9) != true {
		t.Error("Expected distance of 9, got instead ", distance)
	}
}
