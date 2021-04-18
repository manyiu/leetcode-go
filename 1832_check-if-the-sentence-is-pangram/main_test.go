package main

import "testing"

func TestCheckIfPangram(t *testing.T) {
	e1 := checkIfPangram("thequickbrownfoxjumpsoverthelazydog")
	if e1 == false {
		t.Error("Example 1 is not correct.")
	}

	e2 := checkIfPangram("leetcode")
	if e2 == true {
		t.Error("Example 2 is not correct.")
	}
}
