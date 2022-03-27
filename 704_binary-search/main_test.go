package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	e1 := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	if e1 != 4 {
		fmt.Println(e1)
		t.Error("Example 1 is not correct")
	}

	e2 := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	if e2 != -1 {
		fmt.Println(e2)
		t.Error("Example 1 is not correct")
	}
}
