package main

import "testing"

func TestCountConsistentStrings(t *testing.T) {
	e1 := countConsistentStrings("ab", []string{"ad","bd","aaab","baa","badab"})
	if e1 != 2 {
		t.Error("Example 1 is not correct.")
	}

	e2 := countConsistentStrings("abc", []string{"a","b","c","ab","ac","bc","abc"})
	if e2 != 7 {
		t.Error("Example 2 is not correct.")
	}

	e3 := countConsistentStrings("cad", []string{"cc","acd","b","ba","bac","bad","ac","d"})
	if e3 != 4 {
		t.Error("Example 3 is not correct.")
	}
}