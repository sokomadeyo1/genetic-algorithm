package genetic_algorithm

import (
	"math/rand"
	"testing"
)

func Test_sampleIndex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dist      []float64
		rng       rand.Rand
		n_samples int
	}{
		{
			"good_01",
			[]float64{0.2, 0.4, 0.6, 0.8, 1.0},
			*rand.New(rand.NewSource(39)),
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := len(tt.dist)
			for range(tt.n_samples) {
				sample := sampleIndex(tt.dist, tt.rng)
				if sample < 0 || sample > n - 1 {
					t.Errorf("sampleIndex(%d) = %d > %d", n - 1, sample, n - 1)
					break
				}
			}
		})
	}
}
