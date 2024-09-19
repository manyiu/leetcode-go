package removeduplicatesfromsortedarrayii

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}

	for _, test := range tests {
		if result := removeDuplicates(test.nums); result != test.expected {
			t.Errorf("removeDuplicates(%v) = %d, expected = %d", test.nums, result, test.expected)
		}
	}
}
