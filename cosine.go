package similarity

import (
	"math"
)

// Cosine similarity metric finds the normalized dot product of the two attributes.
// By determining the cosine similarity, we would effectively try to find the
// cosine of the angle between the two objects. The cosine of 0° is 1, and it is
// less than 1 for any other angle.

// It is thus a judgement of orientation and not magnitude: two vectors with the same orientation
// have a cosine similarity of 1, two vectors at 90° have a similarity of 0, and two vectors
// diametrically opposed have a similarity of -1, independent of their magnitude.

// Cosine similarity is particularly used in positive space, where the outcome is
// neatly bounded in [0,1]. One of the reasons for the popularity of cosine similarity
// is that it is very efficient to evaluate, especially for sparse vectors.

func CosineSimilarity(coords1, coords2 []float64) (float64, error) {
	var dotProduct float64
	var a float64
	var b float64

	for index, coord := range coords1 {
		dotProduct += coord * coords2[index]

		a += math.Pow(coords1[index], 2)
		b += math.Pow(coords2[index], 2)
	}

	return dotProduct / (math.Sqrt(a) * math.Sqrt(b)), nil
}


func CosineDistance(coords1, coords2 []float64) (float64, error) {
	cosineSim, err := CosineSimilarity(coords1, coords2)
	if err != nil {
		return 0, err
	}

	return (1 - cosineSim), nil
}