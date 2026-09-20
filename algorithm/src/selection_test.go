package genetic_algorithm_test

import (
	"math/rand"
	"testing"

	"github.com/sokomadeyo1/genetic-algorithm/src"
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
			if sum != 1 {
				t.Fatalf("probability does not add up to 1: %g", sum)
			}
		})
	}
}

func TestSampleIndexes(t *testing.T) {
	rng := *rand.New(rand.NewSource(39))
	amount_tests := []struct{
		name string // description of this test case
		// Named input parameters for target function.
		cdf         []float64
		n_samples   int
		n_tests     int
	}{
		{
			"amount_01",
			[]float64{0.2, 0.5, 0.7, 1.0},
			3,
			40,
		},
	}
	for _, tt := range amount_tests {
		t.Run(tt.name, func(t *testing.T) {
			for range tt.n_tests {
				samples := genetic_algorithm.SampleIndexes(tt.cdf, tt.n_samples, rng)
				if len(samples) != tt.n_samples {
					t.Errorf("Invalid amount of samples: %d != %d", len(samples), tt.n_samples)
				}
			}
		})
	}

	repeat_tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cdf         []float64
		n_samples   int
		n_tests     int
	}{
		{
			"repeat_01",
			[]float64{0.2, 0.5, 0.7, 1.0},
			3,
			40,
		},
	}
	for _, tt := range repeat_tests {
		t.Run(tt.name, func(t *testing.T) {
			for range tt.n_tests {
				samples := genetic_algorithm.SampleIndexes(tt.cdf, tt.n_samples, rng)
				for i := range samples {
					for j := i + 1; j < tt.n_samples; j++ {
						if samples[i] == samples[j] {
							t.Errorf("Two similar elements found")
						}
					}
				}
			}
		})
	}
}
