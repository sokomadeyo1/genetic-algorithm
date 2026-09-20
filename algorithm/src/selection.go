package genetic_algorithm

import (
	"math"
	"math/rand"
)

// Evaluate path's length/cost
func Cost(solution Species, task *WGraph) float64 {
	var cost float64 = 0
	n := len(solution)
	for i := range n {
		cost += task.At(solution[i], solution[(i+1)%n])
	}
	return cost
}

func Fitness(solution Species, task *WGraph) float64 {
	return math.Pow(Cost(solution, task), -1)
}

func SoftMax(vals []float64, temperature float64) []float64 {
	var res []float64
	var sum float64 = 0
	for _, val := range vals {
		softmax := math.Exp(val / temperature)
		res = append(res, softmax)
		sum += softmax
	}
	for i := range res {
		res[i] /= sum
	}
	return distFromWeights(res)
}

// Samples k indexes without replacement
func SampleIndexes(cdf []float64, n_samples int, rng rand.Rand) []int {
	n := len(cdf)
	var result []int
	for range n_samples {
		roll := sampleInt(cdf, rng)

		// remove sample
		// TODO: optimize
		var weight float64
		if roll == 0 {
			weight = cdf[roll]
		} else {
			// BUG: index out of range [4] with length 4, see test_SampleIndexes
			weight = cdf[roll] - cdf[roll-1]
		}
		for i := roll; i < n; i++ {
			cdf[i] -= weight
		}
		norm := cdf[n-1]
		for i := range cdf {
			cdf[i] /= norm
		}

		result = append(result, roll)
	}

	return result
}

// Samples k species
func Selection(
	task *WGraph, species []Species, n_samples int, temperature float64, rng rand.Rand,
) []Species {
	// edge case of remain >= len(species)
	if n_samples >= len(species) {
		return species
	}

	var costs []float64
	for _, chromosome := range species {
		costs = append(costs, -Cost(chromosome, task))
	}
	dist := SoftMax(costs, temperature)
	indxs := SampleIndexes(dist, n_samples, rng)
	var result []Species
	for _, i := range indxs {
		result = append(result, species[i])
	}
	return result
}

// Samples a Poisson distributed amount of species
func SelectionPois(
	task *WGraph, species []Species, selection_rate float64, temperature float64, rng rand.Rand,
) []Species {
	n := len(species)
	n_samples := Poisson(selection_rate, n, rng)
	return Selection(task, species, n_samples, temperature, rng)
}
