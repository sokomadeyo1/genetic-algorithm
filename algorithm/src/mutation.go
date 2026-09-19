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
	n := len(species)
	species[i].Finish, species[j].Finish = species[j].Finish, species[i].Finish
	species[(i+1)%n].Start, species[(j+1)%n].Start = species[(j+1)%n].Start, species[(i+1)%n].Start
	return species
}
