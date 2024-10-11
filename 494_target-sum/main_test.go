package targetsum

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
		{[]int{1}, 2, 0},
		{[]int{100}, -200, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findTargetSumWays(tt.nums, tt.target); got != tt.want {
				t.Errorf("FindTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
