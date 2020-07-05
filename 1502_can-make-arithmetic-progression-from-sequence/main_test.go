package main

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	e1 := canMakeArithmeticProgression([]int{3, 5, 1})
	if e1 != true {
		t.Error("Example 1 not correct")
	}

	e2 := canMakeArithmeticProgression([]int{1, 2, 4})
	if e2 != false {
		t.Error("Example 2 not correct")
	}
}
