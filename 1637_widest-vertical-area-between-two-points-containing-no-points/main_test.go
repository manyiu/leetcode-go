package main

import "testing"

func TestMaxWidthOfVerticalArea(t *testing.T) {
	e1 := maxWidthOfVerticalArea([][]int{{8,7},{9,9},{7,4},{9,7}})
	if e1 != 1 {
		t.Error("Example 1 is not correct.")
	}

	e2 := maxWidthOfVerticalArea([][]int{{3,1},{9,0},{1,0},{1,4},{5,3},{8,8}})
	if e2 != 3 {
		t.Error("Example 2 is not correct.")
	}
}