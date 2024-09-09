package searcha2dmatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			3,
			true,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			13,
			false,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			-1,
			false,
		},
		{
			[][]int{},
			0,
			false,
		},
		{
			[][]int{
				{1},
			},
			1,
			true,
		},
		{
			[][]int{
				{1},
			},
			2,
			false,
		},
	}

	for _, test := range tests {
		if got := searchMatrix(test.matrix, test.target); got != test.want {
			t.Errorf("searchMatrix(%v, %v) = %v; want %v\n", test.matrix, test.target, got, test.want)
		}
	}
}
