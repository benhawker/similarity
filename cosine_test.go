package similarity

import (
	"testing"
)

func TestCosineSimilarity(t *testing.T) {
	// Same coords
	a := []float64{1, 1}
	b := []float64{1, 1}
	
	distance, err := CosineSimilarity(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 1) != true {
		t.Error("Expected distance of 1, got instead ", distance)
	}

	// Different coords with 2 dimensions
	a = []float64{1, 2}
	b = []float64{2, 3}

	distance, err = CosineSimilarity(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0.992277877) != true {
		t.Error("Expected distance of 0.992277877, got instead ", distance)
	}

	// Different coords with 3 dimensions
	a = []float64{333, 666, 999}
	b = []float64{999, 333, 666}

	distance, err = CosineSimilarity(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0.7857) != true {
		t.Error("Expected distance of 0.7857, got instead ", distance)
	}
}

func TestCosineDistance(t *testing.T) {
	// Same coords
	a := []float64{1, 1}
	b := []float64{1, 1}
	
	distance, err := CosineDistance(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0) != true {
		t.Error("Expected distance of 0, got instead ", distance)
	}

	// Different coords with 2 dimensions
	a = []float64{1, 2}
	b = []float64{2, 3}

	distance, err = CosineDistance(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0.00772212299999997) != true {
		t.Error("Expected distance of 0.00772212299999997, got instead ", distance)
	}

	// Different coords with 3 dimensions
	a = []float64{333, 666, 999}
	b = []float64{999, 333, 666}

	distance, err = CosineDistance(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0.2142) != true {
		t.Error("Expected distance of 0.2142, got instead ", distance)
	}
}