package laststoneweightii

import "testing"

func TestLastStoneWeightII(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{31, 26, 33, 21, 40}, 5},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1},
	}
	for i, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := lastStoneWeightII(tt.stones); got != tt.want {
				t.Errorf("Test %v: got %v, want %v", i, got, tt.want)
			}
		})
	}
}
