package main

import "testing"

func TestNumberOfMatches(t *testing.T) {
	e1 := numberOfMatches(7)
	if e1 != 6 {
		t.Error("Example 1 is not correct.")
	}

	e2 := numberOfMatches(14)
	if e2 != 13 {
		t.Error("Example 2 is not correct.")
	}
}