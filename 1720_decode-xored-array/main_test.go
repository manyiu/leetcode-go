package main

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	e1 := decode([]int{1,2,3}, 1)
	if reflect.DeepEqual(e1, []int{1,0,2,1}) == false {
		t.Error("Example 1 is not correct.")
	}

	e2 := decode([]int{6,2,7,3}, 4)
	if reflect.DeepEqual(e2, []int{4,2,0,7,4}) == false {
		t.Error("Example 1 is not correct.")
	}
}