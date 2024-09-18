package maximumsumcircularsubarray

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		// {[]int{1, -2, 3, -2}, 3},
		{[]int{5, -3, 5}, 10},
		// {[]int{3, -1, 2, -1}, 4},
		// {[]int{3, -2, 2, -3}, 3},
		// {[]int{-2, -3, -1}, -1},
	}

	for _, test := range tests {
		got := maxSubarraySumCircular(test.nums)
		if got != test.expected {
			t.Errorf("for %v, got: %v, expected: %v", test.nums, got, test.expected)
		}
	}
}
