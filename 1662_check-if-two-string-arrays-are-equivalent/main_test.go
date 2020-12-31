package main

import "testing"

func TestArrayStringsAreEqual(t *testing.T) {
	e1 := arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"})
	if e1 != true {
		t.Error("Example 1 is not correct.")
	}

	e2 := arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"})
	if e2 != false {
		t.Error("Example 2 is not correct.")
	}

	e3 := arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"})
	if e3 != true {
		t.Error("Example 3 is not correct.")
	}
}