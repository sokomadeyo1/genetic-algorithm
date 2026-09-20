package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
	"time"

	genetic_algorithm "github.com/sokomadeyo1/genetic-algorithm/src"
)

const DEFAULT_SIZE = 10
const DEFAULT_ITER = 100
const DEFAULT_POPULATION = 10
const DEFAULT_MUTATION = 0.3
const DEFAULT_CROSSOVER = 0.3
const DEFAULT_GROUPS_P = 0.1
const DEFAULT_TEMPERATURE = 1
const PLACEHOLDER = "PLACEHOLDER"

func main() {
	var n_cities = flag.Int("cities", DEFAULT_SIZE, "number of cities in the TSP")
	var n_iter = flag.Int("iter", DEFAULT_ITER, "number of algorithm iterations")
	var max_population = flag.Int("population", DEFAULT_POPULATION, "the maximum population size after selection")
	var mutation_rate = flag.Float64("mutation", DEFAULT_MUTATION, "mutation probability")
	var crossover_rate = flag.Float64("crossover", DEFAULT_MUTATION, "crossover probability (share of reproducing species)")
	var p_groups = flag.Float64("p", DEFAULT_GROUPS_P, "geometric distribution parameter for crossover groups sizes")
	var temperature = flag.Float64("temp", DEFAULT_TEMPERATURE, "softmax temperature")
	var rng_seed = flag.Int("seed", 39, "random seed")
	var output_file = flag.String("o", PLACEHOLDER, "output file for saving algorithm run data")
	flag.Parse()
	rng := rand.New(rand.NewSource(int64(*rng_seed)))

	problem := genetic_algorithm.Create(*n_cities)

	// Initialize
	var species []genetic_algorithm.Species
	for range *max_population {
		species = append(species, genetic_algorithm.Produce(problem, *rng))
	}

	var simulation [][]genetic_algorithm.Species

	// Main loop
	for gen := range *n_iter {
		// Selection
		reproducing := genetic_algorithm.SelectionPois(problem, species, *crossover_rate, *temperature, *rng)

		// Crossover
		groups := genetic_algorithm.RandomSubsets2(reproducing, *p_groups, *rng)
		for i := range len(groups) {
			child := genetic_algorithm.Crossover(groups[i], problem, *rng)
			species = append(species, child)
		}

		// Mutation
		for i := range len(species) {
			if rng.Float64() <= *mutation_rate {
				species[i] = genetic_algorithm.Mutation(species[i], *rng)
			}
		}

		// Natural selection
		// species = genetic_algorithm.Selection(problem, species, *max_population, *temperature, *rng)
		// TODO: fix sampling and use Selection
		species = species[len(species)-*max_population:]

		// Save data
		simulation = append(simulation, species)

		// Calculate average length
		var costs []float64
		var score float64 = 0
		for _, chromosome := range species {
			costs = append(costs, genetic_algorithm.Cost(chromosome ,problem))
			score += genetic_algorithm.Cost(chromosome, problem)
		}
		slices.Sort(costs)
		score /= float64(len(species))
		fmt.Printf("Generation %4d; path length avg: %g\n", gen+1, score)
		fmt.Printf("                            med: %g\n", costs[len(costs) / 2])
	}

	record := genetic_algorithm.Record{
		City_count:      *n_cities,
		Population:      *max_population,
		Mutation_rate:   *mutation_rate,
		Crossover_rate:  *crossover_rate,
		Crossover_group: *p_groups,
		Seed:            *rng_seed,
		Simulation:      simulation,
	}
	if *output_file == PLACEHOLDER {
		*output_file = fmt.Sprintf("data/simulation_%s.yaml", time.Now().Format(time.DateTime))
	}
	err := genetic_algorithm.YamlWrite(record, *output_file)
	if err != nil {
		panic(err)
	}
}
