package genetic_algorithm

import (
	"math"
	"math/rand"
	"slices"

	"gonum.org/v1/gonum/mat"
)

// Generate a problem
func Create(n int) *WGraph {
	/*
	 * For our problem we'll choose a graph of the following form (denoted using
	 * a matrix):
	 * 0 1 2 1
	 * 1 0 1 2
	 * 2 1 0 1
	 * 1 2 1 0
	 */

	data := make([]float64, n*n)
	for i := range n {
		for j := i; j < n; j++ {
			data[i*n+j] = (float64)(min(j-i, n-(j-i)))
		}
	}
	return (*WGraph)(mat.NewSymDense(n, data))
}

// Generate a random solution using pathfinding algorithm
func Produce(task *WGraph) Species {
	return producer(task, nil)
}

// Solution must be a loop path of length n
func producer(task *WGraph, acc Path) Species {
	n, _ := task.Dims()

	// the current city
	var from int
	if len(acc) == 0 {
		from = 0 // start from the first city
	} else {
		from = acc[len(acc)-1].Finish
	}

	// recursion base
	if len(acc) == n-1 {
		if task.At(from, 0) == math.Inf(1) {
			return nil
		} else {
			return append(acc, Edge{from, 0})
		}
	}

	// cities not to return to
	var visited []int
	for i := range len(acc) {
		visited = append(visited, acc[i].Start)
	}

	var available_ways []int
	for i := range n {
		if from == i {
			continue
		}
		// inf denotes way absence
		if task.At(from, i) != math.Inf(1) && !slices.Contains(visited, i) {
			available_ways = append(available_ways, i)
		}
	}

	for len(available_ways) != 0 {
		next := rand.Intn(len(available_ways))
		new_path := producer(task, append(acc, Edge{from, available_ways[next]}))
		if new_path != nil {
			return new_path
		}
		// delete visited way from available ways
		available_ways[next] = available_ways[len(available_ways)-1]
		available_ways = available_ways[:len(available_ways)-1]
	}
	// no path found
	return nil
}

// Evaluate path's length/cost
func Cost(solution Species, task *WGraph) float64 {
	var cost float64 = 0
	for i := range len(solution) {
		cost += task.At(solution[i].Start, solution[i].Finish)
	}
	return cost
}

func Fitness(solution Species, task *WGraph) float64 {
	return math.Pow(Cost(solution, task), -1)
}

// Samples k species
func Selection(task *WGraph, species []Species, remain int) []Species {
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
	for _ = range remain {
		roll := sampleIndex(dist)
		for slices.Contains(remainingInd, roll) {
			roll = sampleIndex(dist)
		}
		remainingInd = append(remainingInd, roll)
	}

	var remaining []Species
	for i := range remain {
		remaining = append(remaining, species[remainingInd[i]])
	}
	return remaining
}

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

// Random swap
func Mutation(species Species) Species { return nil }

// Unite solutions and generate a new one on a resulting graph
func Crossover(species []Species) Species { return nil }
