package genetic_algorithm

import (
	"math"
	"math/rand"
	"os"

	"gopkg.in/yaml.v3"
)

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

// Samples an integer value from [0, len(dist)-1] based on distribution
func sampleIndex(dist []float64, rng rand.Rand) int {
	p := rng.Float64()
	var i int
	for i = 0; i < len(dist) && p > dist[i]; i++ {
	}
	return i
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Sample Poisson distribution
func Poisson(p float64, n int, rng rand.Rand) int {
	lambda := p * float64(n)
	var weights []float64
	for k := range n {
		weights = append(weights, math.Pow(lambda, float64(k))*math.Exp(-lambda)/float64(factorial(k)))
	}
	dist := distFromWeights(weights)
	return sampleIndex(dist, rng)
}

// Randomly divides a slice into subslices of len >= 2
func RandomSubsets2[T any](items []T, prob float64, rng rand.Rand) [][]T {
	var result [][]T
	if len(items) <= 2 {
		return result
	}

	for i := 0; i < len(items)-1; i++ {
		var subset []T
		subset = append(subset, items[i], items[i+1])
		i++
		// < len - 3 to ensure no subsets of len 1
		for i < len(items)-3 && rng.Float64() < prob {
			subset = append(subset, items[i+1])
			i++
		}
		result = append(result, subset)
	}

	return result
}

func YamlWrite(r Record, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	data, err := yaml.Marshal(r)
	if err != nil {
		return err
	}
	f.Write(data)

	return nil
}
