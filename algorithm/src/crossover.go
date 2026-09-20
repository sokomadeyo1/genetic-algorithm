package genetic_algorithm

import (
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

// Unite solutions and generate a new one on a resulting graph
func Crossover(species []Species, task *WGraph, rng rand.Rand) Species {
	n, _ := task.Dims()
	data := make([]float64, n*n)
	for i := range n {
		for j := i; j < n; j++ {
			data[i*n+j] = math.Inf(1)
		}
	}
	for i := range len(species) {
		for j := range n {
			k := species[i][j]
			l := species[i][(j+1)%n]
			data[k*n+l] = task.At(k, l)
			data[l*n+k] = task.At(k, l)
		}
	}
	subtask := mat.NewSymDense(n, data)
	return Produce(subtask, rng)
}
