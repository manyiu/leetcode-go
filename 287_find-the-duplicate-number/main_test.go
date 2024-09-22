package findtheduplicatenumber

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{1, 3, 4, 2, 2},
			2,
		},
		{
			[]int{3, 1, 3, 4, 2},
			3,
		},
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 1, 2},
			1,
		},
	}

	for _, test := range tests {
		result := findDuplicate(test.nums)
		if result != test.result {
			t.Fatalf("expected %d but got %d", test.result, result)
		}
	}
}
