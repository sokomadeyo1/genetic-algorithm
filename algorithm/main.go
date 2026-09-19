package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/sokomadeyo1/genetic-algorithm/src"
)

const DEFAULT_SIZE = 10
const DEFAULT_ITER = 100
const DEFAULT_POPULATION = 10
const DEFAULT_MUTATION = 0.3
const DEFAULT_CROSSOVER = 0.3
const DEFAULT_GROUPS_P = 0.1

func main() {
	var n_cities = flag.Int("cities", DEFAULT_SIZE, "number of cities in the TSP")
	var n_iter = flag.Int("iter", DEFAULT_ITER, "number of algorithm iterations")
	var max_population = flag.Int("population", DEFAULT_POPULATION, "the maximum population size after selection")
	var mutation_rate = flag.Float64("mutation", DEFAULT_MUTATION, "mutation probability")
	var crossover_rate = flag.Float64("crossover", DEFAULT_MUTATION, "crossover probability (share of reproducing species)")
	var p_groups = flag.Float64("p", DEFAULT_GROUPS_P, "geometric distribution parameter for crossover groups sizes")
	var rng_seed = flag.Int("seed", 39, "random seed")
	flag.Parse()
	// TODO: use this rng in all stochastic functions
	var rng rand.Rand
	rng.Seed(int64(*rng_seed))

	problem := genetic_algorithm.Create(*n_cities)

	for i := range *n_cities {
		for j := range *n_cities {
			fmt.Printf("%g ", problem.At(i, j))
		}
		fmt.Println()
	}

	path := genetic_algorithm.Produce(problem)
	for i := range len(path) {
		fmt.Printf("(%d, %d) ", path[i].Start, path[i].Finish)
	}
	fmt.Println()
	fmt.Println(genetic_algorithm.Cost(path, problem))
}
