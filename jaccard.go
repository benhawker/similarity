package similarity

// The Jaccard similarity measures the similarity between finite sample
// sets and is defined as the cardinality of the intersection of sets 
// divided by the cardinality of the union of the sample sets. 

// Suppose you want to find Jaccard similarity between two sets A and B
// it is the ration of cardinality of A ∩ B and A ∪ B
// func JaccardCoefficient(coords1, coords2 []float64) (float64, error) {
// 	var intersection []float64
// 	var union []float64

// 	for index, coord := range coords1 {
// 		if contains(coords2, coord) && !contains(intersection, coord) {
// 			intersection = append(intersection, coord)
// 		}

// 		union = append(union, coord)
// 	}

// 	for index, coord := range coords2 {
// 		if !contains(intersection, coord) {
// 			union = append(union, coord)
// 		}
// 	}

// 	return len(intersection) / len(union), nil
// }
