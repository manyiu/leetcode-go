package kclosestpointstoorigin

import "testing"

func TestKClosest(t *testing.T) {
	testCases := []struct {
		points   [][]int
		k        int
		expected [][]int
	}{
		{
			points:   [][]int{{1, 3}, {-2, 2}},
			k:        1,
			expected: [][]int{{-2, 2}},
		},
		{
			points:   [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:        2,
			expected: [][]int{{3, 3}, {-2, 4}},
		},
	}

	for _, tc := range testCases {
		if res := kClosest(tc.points, tc.k); !isSame(tc.expected, res) {
			t.Fatalf("expected: %v, got: %v", tc.expected, res)
		}
	}
}

func isSame(i [][]int, res [][]int) bool {
	if len(i) != len(res) {
		return false
	}

	for j := 0; j < len(i); j++ {
		if i[j][0] != res[j][0] || i[j][1] != res[j][1] {
			return false
		}
	}

	return true
}
