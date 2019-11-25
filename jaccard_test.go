package similarity

import (
	"testing"
)

func TestJaccardCoefficient(t *testing.T) {
	a := []float64{1, 2}
	b := []float64{1, 3, 4}
	c := []float64{3, 4}
	
	distance, err := JaccardCoefficient(a, a)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 1.00) != true {
		t.Error("Expected distance of 1.00, got instead ", distance)
	}

	distance, err = JaccardCoefficient(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0.25) != true {
		t.Error("Expected distance of 0.25, got instead ", distance)
	}

	distance, err = JaccardCoefficient(a, c)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0.00) != true {
		t.Error("Expected distance of 0.00, got instead ", distance)
	}

	distance, err = JaccardCoefficient(b, c)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance,  0.6666666666666666) != true {
		t.Error("Expected distance of 0.6666666666666666, got instead ", distance)
	}

	a = []float64{1, 2, 3, 4}
	b = []float64{1, 3, 4}

	distance, err = JaccardCoefficient(a, b)
	if err != nil {
		t.Error(err)
	}
	if almostEqual(distance, 0.75) != true {
		t.Error("Expected distance of 0.75, got instead ", distance)
	}
}