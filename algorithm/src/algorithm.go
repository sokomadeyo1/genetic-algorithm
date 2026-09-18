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
	 * 0 1 2 3
	 * 1 0 1 2
	 * 2 1 0 1
	 * 3 2 1 0
	 */

	data := make([]float64, n*n)
	for i := range n {
		for j := i; j < n; j++ {
			data[i*n+j] = (float64)(j - i)
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
	if from == 0 && len(acc) == n {
		return Species(acc)
	}

	// cities not to return to
	var visited []int
	for i := range len(acc) {
		visited = append(visited, acc[i].Finish)
	}

	var available_ways []int
	for i := 0; i < n; i++ {
		// inf denotes way absence
		if from == i {
			continue
		}
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

// Selects top n species
func Selection(species []Species) []Species { return nil }

// Random swap
func Mutation(species Species) Species { return nil }

// Unite solutions and generate a new one on a resulting graph
func Crossover(species []Species) Species { return nil }
