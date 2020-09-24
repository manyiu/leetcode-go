package main

import "testing"

func TestRestoreString(t *testing.T) {
	e1 := restoreString("codeleet", []int{4,5,6,7,0,2,1,3})
	if e1 != "leetcode" {
		t.Error("Example 1 not correct")
	}

	e2 := restoreString("abc", []int{0,1,2})
	if e2 != "abc" {
		t.Error("Example 2 not correct")
	}

	e3 := restoreString("aiohn", []int{3,1,4,2,0})
	if e3 != "nihao" {
		t.Error("Example 3 not correct")
	}

	e4 := restoreString("aaiougrt", []int{4,0,2,6,7,3,1,5})
	if e4 != "arigatou" {
		t.Error("Example 4 not correct")
	}

	e5 := restoreString("art", []int{1,0,2})
	if e5 != "rat" {
		t.Error("Example 5 not correct")
	}
}