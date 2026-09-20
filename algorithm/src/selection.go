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

// Samples k species
func Selection(task *WGraph, species []Species, remain int, rng rand.Rand) []Species {
	// edge case of remain >= len(species)
	if remain >= len(species) {
		return species
	}

	var weights []float64
	for i := range len(species) {
		w := Fitness(species[i], task)
		weights = append(weights, w)
	}
	dist := distFromWeights(weights)

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
