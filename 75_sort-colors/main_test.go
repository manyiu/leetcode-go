package sortcolors

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TestSortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{2, 0, 2, 1, 1, 0}}, []int{0, 0, 1, 1, 2, 2}},
		{"", args{[]int{2, 0, 1}}, []int{0, 1, 2}},
		{"", args{[]int{0}}, []int{0}},
		{"", args{[]int{1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !equal(tt.args.nums, tt.want) {
				t.Errorf("SortColors() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
