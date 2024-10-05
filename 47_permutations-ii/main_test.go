package permutationsii

import "testing"

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

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{
			name:   "test1",
			input:  []int{1, 1, 2},
			expect: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name:  "test2",
			input: []int{1, 2, 3},
			expect: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
				{2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name:   "test3",
			input:  []int{0, 1},
			expect: [][]int{{0, 1}, {1, 0}},
		},
		{
			name:   "test4",
			input:  []int{1},
			expect: [][]int{{1}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := permuteUnique(test.input)
			if !equal(test.expect, got) {
				t.Errorf("expect %v, got %v", test.expect, got)
			}
		})
	}
}
