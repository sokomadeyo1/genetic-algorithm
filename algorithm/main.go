package main

import (
	"flag"
	"fmt"
	"math/rand"

	genetic_algorithm "github.com/sokomadeyo1/genetic-algorithm/src"
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
	rng := rand.New(rand.NewSource(int64(*rng_seed)))

	problem := genetic_algorithm.Create(*n_cities)

	// Initialize
	var species []genetic_algorithm.Species
	for range *max_population {
		species = append(species, genetic_algorithm.Produce(problem, *rng))
	}

	// Main loop
	for gen := range *n_iter {
		// Selection
		species = genetic_algorithm.Selection(problem, species, *max_population, *rng)

		// Crossover
		var reproducing []genetic_algorithm.Species
		for i := range len(species) {
			if rng.Float64() <= *crossover_rate {
				reproducing = append(reproducing, species[i])
			}
		}
		groups := genetic_algorithm.RandomSubsets2(reproducing, *p_groups, *rng)
		for i := range len(groups) {
			child := genetic_algorithm.Crossover(groups[i], problem, *rng)
			if child != nil {
				species = append(species, genetic_algorithm.Crossover(groups[i], problem, *rng))
			} else {
				fmt.Println("Nil child")
			}
		}

		// Mutation
		for i := range len(species) {
			if rng.Float64() <= *mutation_rate {
				species[i] = genetic_algorithm.Mutation(species[i], *rng)
			}
		}

		// Calculate average length
		var score float64 = 0
		for i := range len(species) {
			score += genetic_algorithm.Cost(species[i], problem)
		}
		score /= float64(len(species))
		fmt.Printf("Generation %4d; average path length: %g\n", gen+1, score)
	}
}
