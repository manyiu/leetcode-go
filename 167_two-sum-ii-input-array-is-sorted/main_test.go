package twosumiiinputarrayissorted

import (
	"testing"
)

func equal(a, b []int) bool {
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

func TestTwoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	type test struct {
		name string
		args args
		want []int
	}

	tests := []test{
		{
			name: "Example 1",
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			name: "Example 2",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			name: "Example 3",
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !equal(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
