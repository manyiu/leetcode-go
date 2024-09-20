package findpivotindex

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"case1", []int{1, 7, 3, 6, 5, 6}, 3},
		{"case2", []int{1, 2, 3}, -1},
		{"case3", []int{2, 1, -1}, 0},
		{"case4", []int{0, 0, 0, 0, 0, 0}, 0},
		{"case5", []int{1, 2, 3, 4, 5, 6}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.nums); got != tt.want {
				t.Errorf("PivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
