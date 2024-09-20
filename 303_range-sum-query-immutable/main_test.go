package rangesumqueryimmutable

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{"case1", 0, 2, 1},
		{"case2", 2, 5, -1},
		{"case3", 0, 5, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obj.SumRange(tt.left, tt.right); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
