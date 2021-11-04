package main

import (
	"fmt"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	e1 := minCostClimbingStairs([]int{10, 15, 20})
	if e1 != 15 {
		fmt.Println(e1)
		t.Error("Example 1 is not correct.")
	}

	e2 := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	if e2 != 6 {
		fmt.Println(e2)
		t.Error("Example 2 is not correct.")
	}
}
