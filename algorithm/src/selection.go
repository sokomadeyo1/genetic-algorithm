package genetic_algorithm

import (
	"math"
	"math/rand"
	"slices"
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

// Samples k species
func Selection(
	task *WGraph, species []Species, remain int, temperature float64, rng rand.Rand,
) []Species {
	// edge case of remain >= len(species)
	if remain >= len(species) {
		return species
	}

	var costs []float64
	for _, chromosome := range species {
		costs = append(costs, -Cost(chromosome, task))
	}
	dist := SoftMax(costs, temperature)

	// we have to do non-return sampling
	var remainingInd []int
	// NOTE: might be more computationally efficient to roll for death instead of survival
	for range remain {
		roll := sampleIndex(dist, rng)
		// NOTE: might lead to redundant rerolls
		for slices.Contains(remainingInd, roll) {
			roll = sampleIndex(dist, rng)
		}
		remainingInd = append(remainingInd, roll)
	}

	var remaining []Species
	for i := range remain {
		remaining = append(remaining, species[remainingInd[i]])
	}
	return remaining
}

// Samples a Poisson distributed amount of species
func SelectionPois(
	task *WGraph, species []Species, selection_rate float64, temperature float64, rng rand.Rand,
) []Species {
	n := len(species)
	n_samples := Poisson(selection_rate, n, rng)

	var costs []float64
	for _, chromosome := range species {
		costs = append(costs, -Cost(chromosome, task))
	}
	dist := SoftMax(costs, temperature)

	// we have to do non-return sampling
	var remainingInd []int
	// NOTE: might be more computationally efficient to roll for death instead of survival
	for range n_samples {
		roll := sampleIndex(dist, rng)
		// NOTE: might lead to redundant rerolls
		for slices.Contains(remainingInd, roll) {
			roll = sampleIndex(dist, rng)
		}
		remainingInd = append(remainingInd, roll)
	}

	var remaining []Species
	for _, i := range remainingInd {
		remaining = append(remaining, species[i])
	}
	return remaining
}
