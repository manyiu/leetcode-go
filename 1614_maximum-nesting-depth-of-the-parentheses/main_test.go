package main

import "testing"

func TestMaxDepth(t *testing.T) {
	e1 := maxDepth("(1+(2*3)+((8)/4))+1")
	if e1 != 3 {
		t.Error("Example 1 is not correct.")
	}

	e2 := maxDepth("(1)+((2))+(((3)))")
	if e2 != 3 {
		t.Error("Example 1 is not correct.")
	}

	e3 := maxDepth("1+(2*3)/(2-1)")
	if e3 != 1 {
		t.Error("Example 1 is not correct.")
	}

	e4 := maxDepth("1")
	if e4 != 0 {
		t.Error("Example 1 is not correct.")
	}
}
