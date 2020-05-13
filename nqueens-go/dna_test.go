package main

import (
	"testing"
)

func TestDNA_Fitness(t *testing.T) {
	tests := []struct {
		name      string
		positions []int
		want      float64
	}{
		{
			name:      "all horizontal",
			positions: []int{0, 0, 0, 0, 0, 0, 0, 0},
			want:      0,
		},
		{
			name:      "all horizontal mid row",
			positions: []int{4, 4, 4, 4, 4, 4, 4, 4},
			want:      0,
		},
		{
			name:      "all diagonal",
			positions: []int{0, 1, 2, 3, 4, 5, 6, 7},
			want:      0,
		},
		{
			name:      "quarter conflicted",
			positions: []int{0, 6, 4, 7, 5, 6, 5, 6},
			want:      0.75,
		},
		{
			name:      "solved",
			positions: []int{4, 2, 0, 6, 1, 7, 5, 3},
			want:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dna := &DNA{
				Positions: tt.positions,
			}
			if got := dna.Fitness(); got != tt.want {
				t.Errorf("Fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}
