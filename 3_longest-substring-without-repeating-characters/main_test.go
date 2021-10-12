package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	e1 := lengthOfLongestSubstring("abcabcbb")
	if e1 != 3 {
		t.Error("Example 1 is not correct.")
	}

	e2 := lengthOfLongestSubstring("bbbbb")
	if e2 != 1 {
		t.Error("Example 2 is not correct.")
	}

	e3 := lengthOfLongestSubstring("pwwkew")
	if e3 != 3 {
		t.Error("Example 3 is not correct.")
	}

	e4 := lengthOfLongestSubstring("")
	if e4 != 0 {
		t.Error("Example 4 is not correct.")
	}
}
