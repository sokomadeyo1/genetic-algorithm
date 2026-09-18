package main

import (
	"flag"
	"fmt"

	"github.com/sokomadeyo1/genetic-algorithm/src"
)

const DEFAULT_SIZE = 10
const DEFAULT_ITER = 10

func main() {
	var n_cities = flag.Int("n_cities", DEFAULT_SIZE, "number of cities in the TSP")
	// var n_iter = flag.Int("n_iter", DEFAULT_ITER, "number of algorithm iterations")
	flag.Parse()

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
