package genetic_algorithm

import "testing"

func Test_swap(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		species Species
		i       int
		j       int
		want    Species
	}{
		{
			"good_01",
			Species{Edge{0, 1}, Edge{1, 2}, Edge{2, 0}},
			1,
			2,
			Species{Edge{2, 1}, Edge{1, 0}, Edge{0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swap(tt.species, tt.i, tt.j)
			pass := true
			for i := range len(got) {
				if got[i] != tt.want[i] {
					pass = false
					break
				}
			}
			if !pass {
				t.Errorf("swap() = %v, want %v", got, tt.want)
			}
		})
	}
}
