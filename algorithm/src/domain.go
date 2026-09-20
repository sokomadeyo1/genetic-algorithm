package genetic_algorithm

import "gonum.org/v1/gonum/mat"

type Species = Path
type Path = []int

// Weighted undirected graph
type WGraph = mat.SymDense

// Record for saving data in yaml
type Record struct {
	City_count      int
	Population      int
	Mutation_rate   float64
	Crossover_rate  float64
	Crossover_group float64
	Seed            int
	Mean_length     float64
	Median_length   int
	Simulation      [][]Species
}
