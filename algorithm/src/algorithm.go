package genetic_algorithm

import (
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

// Generate a random solution
func Generate(task WGraph) Species { return nil }

func Score(solution Species, task WGraph) int { return 0 }

// Selects top n species
func Selection(species []Species) []Species { return nil }

// Random swap
func Mutation(species Species) Species { return nil }

// Unite solutions and generate a new one on a resulting graph
func Crossover(species []Species) Species { return nil }
