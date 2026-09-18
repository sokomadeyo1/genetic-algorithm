package genetic_algorithm

import "gonum.org/v1/gonum/mat"

type Species = Path
type Path = []Edge
type Edge struct {
	Start int
	Finish int
}
// Weighted undirected graph
type WGraph = mat.SymDense
