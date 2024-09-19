package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}

	for _, test := range tests {
		if result := maxArea(test.height); result != test.expected {
			t.Errorf("maxArea(%v) = %d, expected = %d", test.height, result, test.expected)
		}
	}
}
