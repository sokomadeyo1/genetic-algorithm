package genetic_algorithm_test

import (
	"github.com/sokomadeyo1/genetic-algorithm/src"
	"testing"
)

func TestSoftMax(t *testing.T) {
	norm_tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		vals        []float64
		temperature float64
	}{
		{
			"edge_01",
			[]float64{1},
			1,
		},
		{
			"edge_02",
			[]float64{1, 2},
			1,
		},
		{
			"good_01",
			[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1,
		},
		{
			"good_02",
			[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			1,
		},
	}
	for _, tt := range norm_tests {
		t.Run(tt.name, func(t *testing.T) {
			got := genetic_algorithm.SoftMax(tt.vals, tt.temperature)
			for _, p := range got {
				if p < 0 || p > 1 {
					t.Fatalf("invalid probability: %g", p)
				}
			}
			sum := got[len(got)-1]
			if sum  != 1 {
				t.Fatalf("probability does not add up to 1: %g", sum)
			}
		})
	}
}
