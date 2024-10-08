package findcriticalandpseudocriticaledgesinminimumspanningtree

import "testing"

func compare(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func TestFindCriticalAndPseudoCriticalEdges(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		edges  [][]int
		expect [][]int
	}{
		{
			name: "test 1",
			n:    5,
			edges: [][]int{
				{0, 1, 1},
				{1, 2, 1},
				{2, 3, 2},
				{0, 3, 2},
				{0, 4, 3},
				{3, 4, 3},
				{1, 4, 6},
			},
			expect: [][]int{
				{0, 1},
				{2, 3, 4, 5},
			},
		},
		{
			name: "test 2",
			n:    4,
			edges: [][]int{
				{0, 1, 1},
				{1, 2, 1},
				{2, 3, 1},
				{0, 3, 1},
			},
			expect: [][]int{
				{},
				{0, 1, 2, 3},
			},
		},
		{
			name: "test 3",
			n:    6,
			edges: [][]int{
				{0, 1, 1},
				{1, 2, 1},
				{0, 2, 1},
				{2, 3, 4},
				{3, 4, 2},
				{3, 5, 2},
				{4, 5, 2},
			},
			expect: [][]int{
				{3},
				{0, 1, 2, 4, 5, 6},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findCriticalAndPseudoCriticalEdges(test.n, test.edges)
			if !compare(got, test.expect) {
				t.Errorf("expect %v, got %v", test.expect, got)
			}
		})
	}
}
