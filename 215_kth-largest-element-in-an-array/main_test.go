package kthlargestelementinanarray

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{7, 6, 5, 4, 3, 2, 1}, 5, 3},
	}
	for _, tt := range tests {
		if got := findKthLargest(tt.nums, tt.k); got != tt.want {
			t.Errorf("FindKthLargest(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
