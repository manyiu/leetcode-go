package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("search(%v, %d)", test.nums, test.target)
		got := search(test.nums, test.target)
		if got != test.want {
			t.Fatalf("%s = %d, want %d", descr, got, test.want)
		}
	}
}
