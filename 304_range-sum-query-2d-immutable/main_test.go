package rangesumquery2dimmutable

import "testing"

func TestNumMatrix_SumRegion(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)
	tests := []struct {
		name string
		row1 int
		col1 int
		row2 int
		col2 int
		want int
	}{
		{"case1", 2, 1, 4, 3, 8},
		{"case2", 1, 1, 2, 2, 11},
		{"case3", 1, 2, 2, 4, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obj.SumRegion(tt.row1, tt.col1, tt.row2, tt.col2); got != tt.want {
				t.Errorf("NumMatrix.SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

// matrix := [][]int{
// 	{3, 0, 1, 4, 2},
// 	{5, 6, 3, 2, 1},
// 	{1, 2, 0, 1, 5},
// 	{4, 1, 0, 1, 7},
// 	{1, 0, 3, 0, 5},
// }

// prefixSum := [][]int{
// 	{3, 3, 4, 8, 10},
// 	{8, 14, 18, 24, 27},
// 	{9, 17, 21, 28, 36},
// 	{13, 22, 26, 34, 49},
// 	{14, 23, 30, 38, 58},
// }
