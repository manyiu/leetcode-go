package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, want: 9},
		{height: []int{4, 2, 3}, want: 1},
		{height: []int{4, 2, 3, 1}, want: 1},
		{height: []int{4, 2, 3, 1, 2}, want: 2},
		{height: []int{4, 2, 3, 1, 2, 1}, want: 2},
		{height: []int{4, 2, 3, 1, 2, 1, 3}, want: 6},
		{height: []int{4, 2, 3, 1, 2, 1, 3, 1}, want: 6},
		{height: []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}, want: 23},
	}

	for _, tt := range tests {
		if got := trap(tt.height); got != tt.want {
			t.Errorf("trap(%v) = %v, want %v", tt.height, got, tt.want)
		}
	}
}
