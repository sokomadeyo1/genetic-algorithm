package genetic_algorithm_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/sokomadeyo1/genetic-algorithm/src"
	"gonum.org/v1/gonum/mat"
)

func TestProduce(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		task *genetic_algorithm.WGraph
	}{
		{
			"good_01",
			mat.NewSymDense(3, []float64{0, 1, 1, 1, 0, 1, 1, 1, 0}),
		},
		{
			"good_02",
			mat.NewSymDense(4, []float64{0, 1, 2, 1, 1, 0, 1, 2, 2, 1, 0, 1, 1, 2, 1, 0}),
		},
	}
	rng := rand.New(rand.NewSource(39))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := genetic_algorithm.Produce(tt.task, *rng)
			var vals []int
			for _, item := range got {
				if slices.Contains(vals, item) {
					t.Errorf("%v contains %d several times", got, item)
				} else {
					vals = append(vals, item)
				}
			}
		})
	}
}
