package mergeintervals

import (
	"testing"
)

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if len(v) != len(b[i]) {
			return false
		}

		for j, v2 := range v {
			if v2 != b[i][j] {
				return false
			}
		}
	}

	return true
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{
			name: "test1",
			input: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			expect: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "test2",
			input: [][]int{
				{1, 4},
				{4, 5},
			},
			expect: [][]int{
				{1, 5},
			},
		},
		{
			name: "test3",
			input: [][]int{
				{1, 4},
				{0, 4},
			},
			expect: [][]int{
				{0, 4},
			},
		},
		{
			name: "test4",
			input: [][]int{
				{1, 4},
				{0, 0},
			},
			expect: [][]int{
				{0, 0},
				{1, 4},
			},
		},
		{
			name: "test5",
			input: [][]int{
				{1, 4},
				{0, 2},
				{3, 5},
			},
			expect: [][]int{
				{0, 5},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := merge(test.input)
			if !equal(test.expect, got) {
				t.Errorf("expect %v, got %v", test.expect, got)
			}
		})
	}
}
