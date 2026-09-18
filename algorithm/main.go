package main

import (
	"flag"
	"fmt"
)

const DEFAULT_SIZE = 10
const DEFAULT_ITER = 10

func main() {
	var n_cities = flag.Int("n_cities", DEFAULT_SIZE, "number of cities in the TSP")
	var n_iter = flag.Int("n_iter", DEFAULT_ITER, "number of algorithm iterations")
	flag.Parse()

	fmt.Printf("%d %d", *n_cities, *n_iter)
}
