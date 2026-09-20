package genetic_algorithm

import "math/rand"

// Random swap
func Mutation(species Species, rng rand.Rand) Species {
	n := len(species)
	i := rng.Intn(n)
	j := rng.Intn(n)
	for j == i {
		j = rng.Intn(n)
	}

	return swap(species, i, j)
}

func swap(species Species, i, j int) Species {
	species[i], species[j] = species[j], species[i]
	return species
}
