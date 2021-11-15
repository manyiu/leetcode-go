package main

import (
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	e1 := buildArray([]int{0, 2, 1, 5, 3, 4})
	if !reflect.DeepEqual(e1, []int{0, 1, 2, 4, 5, 3}) {
		t.Error("Example 1 is not correct")
	}

	e2 := buildArray([]int{5, 0, 1, 2, 3, 4})
	if !reflect.DeepEqual(e2, []int{4, 5, 0, 1, 2, 3}) {
		t.Error("Example 2 is not correct")
	}
}
