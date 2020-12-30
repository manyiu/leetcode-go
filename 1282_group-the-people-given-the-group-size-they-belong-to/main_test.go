package main

import (
	"reflect"
	"sort"
	"testing"
)

func groupedToUngrouped(grouped [][]int) []int {
	r := []int{}
	
	for _, v := range grouped {
		s := len(v)
		t := make([]int, s)

		for i := 0; i < s; i ++ {
			t[i] = s
		}

		r = append(r, t...)
	}

	return r
}

func TestGroupThePeople(t *testing.T) {
	i1 := []int{3,3,3,3,3,1,3}
	o1 := groupedToUngrouped(groupThePeople(i1))
	sort.Ints(i1)
	sort.Ints(o1)

	if reflect.DeepEqual(i1, o1) == false {
		t.Error("Example 1 is not correct")
	}

	i2 := []int{2,1,3,3,3,2}
	o2 := groupedToUngrouped(groupThePeople(i2))
	sort.Ints(i2)
	sort.Ints(o2)

	if reflect.DeepEqual(i2, o2) == false {
		t.Error("Example 2 is not correct")
	}
}