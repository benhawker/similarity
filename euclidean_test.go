package similarity

import (
	"testing"
)

func TestEuclidean(t *testing.T) {
	// Same coords
	a := []float64{1, 1}
	b := []float64{1, 1}
	
	distance, err := Euclidean(a, b)
	if err != nil {
		t.Error(err)
	}
	if distance != 0 {
		t.Error("Expected distance of 0, got instead ", distance)
	}

	// Different coords with 2 dimensions
	a = []float64{1, 2}
	b = []float64{3, 4}

	distance, err = Euclidean(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 2.828) != true {
		t.Error("Expected distance of 2.828, got instead ", distance)
	}

	// Different coords with 3 dimensions
	a = []float64{1, 2, 3}
	b = []float64{4, 5, 6}

	distance, err = Euclidean(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 5.196) != true {
		t.Error("Expected distance of 5.196, got instead ", distance)
	}
}
