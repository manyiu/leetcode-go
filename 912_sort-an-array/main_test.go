package sortanarray

import "testing"

func TestSortArray(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
		{[]int{5, 1, 1, 2, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 1, 1, 2, 5}},
		{[]int{5, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 5}},
	}
	for _, test := range tests {
		got := sortArray(test.nums)
		if len(got) != len(test.want) {
			t.Errorf("sortArray(%v) = %v; want %v\n", test.nums, got, test.want)
			continue
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("sortArray(%v) = %v; want %v\n", test.nums, got, test.want)
				break
			}
		}
	}
}
