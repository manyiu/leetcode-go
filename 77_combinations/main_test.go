package combinations

import "testing"

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aa := range a {
		if len(aa) != len(b[i]) {
			return false
		}
		for j, v := range aa {
			if v != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestCombine(t *testing.T) {
	tests := []struct {
		n, k int
		want [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{1, 1, [][]int{{1}}},
		{1, 0, [][]int{{}}},
		{4, 3, [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}},
		{5, 4, [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}}},
	}
	for _, test := range tests {
		if got := combine(test.n, test.k); !equal(got, test.want) {
			t.Errorf("combine(%v, %v) = %v; want %v", test.n, test.k, got, test.want)
		}
	}
}
