package sortitemsbygroupsrespectingdependencies

import "testing"

func isSameSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSortItems(t *testing.T) {
	tests := []struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
		want        []int
	}{
		{
			8,
			2,
			[]int{-1, -1, 1, 0, 0, 1, 0, -1},
			[][]int{{}, {6}, {5}, {6}, {3}, {}, {}, {}},
			[]int{6, 3, 4, 5, 2, 0, 7, 1},
		},
		{
			8,
			2,
			[]int{-1, -1, 1, 0, 0, 1, 0, -1},
			[][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}},
			[]int{},
		},
		{
			5,
			5,
			[]int{2, 0, -1, 3, 0},
			[][]int{{2, 1, 3}, {2, 4}, {}, {}, {}},
			[]int{2, 1, 3, 4, 0},
		},
		{
			4,
			1,
			[]int{-1, 0, 0, -1},
			[][]int{{}, {0}, {1, 3}, {2}},
			[]int{},
		},
		{
			5,
			5,
			[]int{2, 0, -1, 3, 0},
			[][]int{{2, 1, 3}, {2, 4}, {}, {}, {}},
			[]int{3, 2, 4, 1, 0},
		},
	}
	for i, test := range tests {
		if got := sortItems(test.n, test.m, test.group, test.beforeItems); !isSameSlice(got, test.want) {
			t.Errorf("Test %v: sortItems(%v, %v, %v, %v) = %v, want %v", i, test.n, test.m, test.group, test.beforeItems, got, test.want)
		}
	}
}
