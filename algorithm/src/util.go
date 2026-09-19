package genetic_algorithm

import "math/rand"

// Converts weights to distribution intervals on [0, 1]
func distFromWeights(weights []float64) []float64 {
	for i := 1; i < len(weights); i++ {
		weights[i] += weights[i-1]
	}
	total := weights[len(weights)-1]
	for i := range len(weights) {
		weights[i] /= total
	}
	return weights
}

// Samples an integer value from [1, len(dist)] based on distribution
func sampleIndex(dist []float64) int {
	p := rand.Float64()
	var i int
	for i = 0; i < len(dist) && p < dist[i]; i++ {
		if p >= dist[i] {
			break
		}
	}
	return i
}

// Randomly divides a slice into subslices of len >= 2
func RandomSubsets2[T any](items []T, prob float64) [][]T {
	var result [][]T
	if len(items) <= 2 {
		return result
	}

	for i := 0; i < len(items)-1; i++ {
		var subset []T
		subset = append(subset, items[i], items[i+1])
		i++
		// < len - 3 to ensure no subsets of len 1
		for i < len(items)-3 && rand.Float64() < prob {
			subset = append(subset, items[i+1])
			i++
		}
		result = append(result, subset)
	}

	return result
}
