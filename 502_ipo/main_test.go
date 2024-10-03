package ipo

import "testing"

func TestFindMaximizedCapital(t *testing.T) {
	tests := []struct {
		k       int
		w       int
		profits []int
		capital []int
		want    int
	}{
		{
			k:       2,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{0, 1, 1},
			want:    4,
		},
		{
			k:       3,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{0, 1, 2},
			want:    6,
		},
	}

	for _, tt := range tests {
		if got := findMaximizedCapital(tt.k, tt.w, tt.profits, tt.capital); got != tt.want {
			t.Errorf("findMaximizedCapital(%v, %v, %v, %v) = %v, want %v", tt.k, tt.w, tt.profits, tt.capital, got, tt.want)
		}
	}
}
