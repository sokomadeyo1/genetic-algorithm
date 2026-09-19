package genetic_algorithm

import "math/rand"

// Random swap
func Mutation(species Species) Species {
	n := len(species)
	i := rand.Intn(n)
	j := rand.Intn(n)
	for j == i {
		j = rand.Intn(n)
	}

	species[i].Finish, species[j].Finish = species[j].Finish, species[i].Finish
	species[(i+1)%n].Start, species[(j+1)%n].Start = species[(j+1)%n].Start, species[(i+1)%n].Start

	return species
}

