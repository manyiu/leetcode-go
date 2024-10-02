package slidingwindowmedian

import "testing"

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMedianSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []float64
	}{
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []float64{1, -1, -1, 3, 5, 6},
		},
		{
			nums: []int{1, 2, 3, 4, 2, 3, 1, 4, 2},
			k:    3,
			want: []float64{2, 3, 3, 3, 2, 3, 2},
		},
	}
	for _, tt := range tests {
		if got := medianSlidingWindow(tt.nums, tt.k); !equal(got, tt.want) {
			t.Errorf("medianSlidingWindow(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
