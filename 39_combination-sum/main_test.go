package combinationsum

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

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expect     [][]int
	}{
		{
			name:       "test1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expect: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			name:       "test2",
			candidates: []int{2, 3, 5},
			target:     8,
			expect: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name:       "test3",
			candidates: []int{2},
			target:     1,
			expect:     [][]int{},
		},
		{
			name:       "test4",
			candidates: []int{1},
			target:     1,
			expect: [][]int{
				{1},
			},
		},
		{
			name:       "test5",
			candidates: []int{1},
			target:     2,
			expect: [][]int{
				{1, 1},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := combinationSum(test.candidates, test.target)
			if !equal(got, test.expect) {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
