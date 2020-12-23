package main

import "testing"

func TestMaximumWealth(t *testing.T) {
	e1 := maximumWealth([][]int{{1,2,3},{3,2,1}})
	if e1 != 6 {
		t.Error("Example 1 not correct")
	}

	e2 := maximumWealth([][]int{{1,5},{7,3}, {3,5}})
	if e2 != 10 {
		t.Error("Example 2 not correct")
	}

	e3 := maximumWealth([][]int{{2,8,7},{7,1,3},{1,9,5}})
	if e3 != 17 {
		t.Error("Example 3 not correct")
	}
}