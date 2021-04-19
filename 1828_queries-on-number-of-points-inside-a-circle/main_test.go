package main

import (
	"reflect"
	"testing"
)

func TestCountPoints(t *testing.T) {
	e1 := countPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}})
	if reflect.DeepEqual(e1, []int{3, 2, 2}) == false {
		t.Error("Example 1 is not correct.")
	}

	e2 := countPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}})
	if reflect.DeepEqual(e2, []int{2, 3, 2, 4}) == false {
		t.Error("Example 2 is not correct.")
	}
}
