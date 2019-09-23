package similarity

import (
	"testing"
)

func TestJaccardCoefficient(t *testing.T) {
	// Same coords
	a := []float64{1, 1}
	b := []float64{1, 1}
	
	distance, err := JaccardCoefficient(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 1) != true {
		t.Error("Expected distance of 1, got instead ", distance)
	}

	// Different coords with 2 dimensions
	a = []float64{1, 2}
	b = []float64{2, 3}

	distance, err = JaccardCoefficient(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0.992277877) != true {
		t.Error("Expected distance of 0.992277877, got instead ", distance)
	}
}