package genetic_algorithm

import (
	"math/rand"
	"slices"
	"testing"
)

func Test_sampleIndex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dist      []float64
		n_samples int
	}{
		{
			"domain_01",
			[]float64{0.2, 0.4, 0.6, 0.8, 1.0},
			100,
		},
		{
			"domain_edge_01",
			[]float64{1.0},
			10,
		},
	}
	rng := *rand.New(rand.NewSource(39))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := len(tt.dist)
			for range tt.n_samples {
				sample := sampleInt(tt.dist, rng)
				if sample < 0 {
					t.Errorf("sampleIndex(%d) = %d < 0", n-1, sample)
					break
				} else if sample > n-1 {
					t.Errorf("sampleIndex(%d) = %d > %d", n-1, sample, n-1)
					break
				}
			}
		})
	}

	// test diversity
	tests = []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dist      []float64
		n_samples int
	}{
		{
			"diversity_01",
			[]float64{0.2, 0.4, 0.6, 0.8, 1.0},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := len(tt.dist)
			var target []int
			for i := range n {
				target = append(target, i)
			}
			for range tt.n_samples {
				sample := sampleInt(tt.dist, rng)
				if slices.Contains(target, sample) {
					target = slices.DeleteFunc(target, func(i int) bool { return i == sample })
				}
				if len(target) == 0 {
					break
				}
			}
			if len(target) != 0 {
				t.Errorf("sampleIndex(%d): not all values sampled: [", n)
				for _, item := range target {
					t.Error(item)
				}
				t.Error("]\n")
			}
		})
	}
}

func TestPoisson(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		p float64
		n int
	}{
		{
			"domain_01",
			0.5,
			4,
		},
		{
			"domain_edge_01",
			0.5,
			1,
		},
	}
	rng := *rand.New(rand.NewSource(39))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for range 100 * tt.n {
				sample := Poisson(tt.p, tt.n, rng)
				if sample < 0 {
					t.Errorf("Poisson(n=%d) = %d < 0", tt.n, sample)
					break
				} else if sample > tt.n-1 {
					t.Errorf("Poisson(n=%d) = %d > %d", tt.n, sample, tt.n-1)
					break
				}
			}
		})
	}

	// test diversity
	diversity_tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		p         float64
		n         int
		n_samples int
	}{
		{
			"diversity_01",
			0.5,
			10,
			1000,
		},
		{
			"diversity_02",
			0.5,
			20,
			100000,
		},
		{
			"diversity_edge_01",
			0.5,
			1,
			100,
		},
		{
			"diversity_edge_02",
			0.5,
			2,
			100,
		},
	}
	for _, tt := range diversity_tests {
		t.Run(tt.name, func(t *testing.T) {
			var target []int
			for i := range tt.n {
				target = append(target, i)
			}
			for range tt.n_samples {
				sample := Poisson(tt.p, tt.n, rng)
				if slices.Contains(target, sample) {
					target = slices.DeleteFunc(target, func(i int) bool { return i == sample })
				}
				if len(target) == 0 {
					break
				}
			}
			if len(target) != 0 {
				t.Errorf("Poisson(n=%d): not all values sampled: [", tt.n)
				for _, item := range target {
					t.Error(item)
				}
				t.Error("]\n")
			}
		})
	}
}
