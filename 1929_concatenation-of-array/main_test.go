package main

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	e1 := getConcatenation([]int{1, 2, 1})
	if !reflect.DeepEqual(e1, []int{1, 2, 1, 1, 2, 1}) {
		t.Error("Example 1 is not correct")
	}

	e2 := getConcatenation([]int{1, 3, 2, 1})
	if !reflect.DeepEqual(e2, []int{1, 3, 2, 1, 1, 3, 2, 1}) {
		t.Error("Example 1 is not correct")
	}
}
