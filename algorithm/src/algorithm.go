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
