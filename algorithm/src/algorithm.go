package genetic_algorithm

// Generate a random solution
func generate(task WGraph) Species { return nil }

func score(solution Species, task WGraph) int { return 0 }

// Selects top n species
func selection(species []Species) []Species { return nil }

// Random swap
func mutation(species Species) Species { return nil }

// Unite solutions and generate a new one on a resulting graph
func crossover(species []Species) Species { return nil }
