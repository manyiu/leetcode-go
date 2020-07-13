package main

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	e1 := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	if e1 != 4 {
		t.Error("Example 1 not correct")
	}

	e2 := numIdenticalPairs([]int{1, 1, 1, 1})
	if e2 != 6 {
		t.Error("Example 2 not correct")
	}

	e3 := numIdenticalPairs([]int{1, 2, 3})
	if e3 != 0 {
		t.Error("Example 3 not correct")
	}
}
