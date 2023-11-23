package removeduplicatesfromsortedarray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var cases = []struct {
		input    []int
		expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tt := range cases {
		actual := removeDuplicates(tt.input)
		if actual != tt.expected {
			t.Errorf("removeDuplicates(%v): expected %d, actual %d", tt.input, tt.expected, actual)
		}
	}
}
