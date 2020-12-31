package main

import "testing"

func TestSubrectangleQueries(t *testing.T) {
	r1 := Constructor([][]int{{1,2,1},{4,3,4},{3,2,1},{1,1,1}})
	
	e1a := r1.GetValue(0,2)
	if e1a != 1 {
		t.Error("Example 1a is not correct.")
	}

	r1.UpdateSubrectangle(0, 0, 3, 2, 5)

	e1b := r1.GetValue(0,2)
	if e1b != 5 {
		t.Error("Example 1b is not correct.")
	}

	e1c := r1.GetValue(3,1)
	if e1c != 5 {
		t.Error("Example 1c is not correct.")
	}

	r1.UpdateSubrectangle(3, 0, 3, 2, 10)

	e1d := r1.GetValue(3,1)
	if e1d != 10 {
		t.Error("Example 1d is not correct.")
	}

	e1e := r1.GetValue(0,2)
	if e1e != 5 {
		t.Error("Example 1e is not correct.")
	}

	r2 := Constructor([][]int{{1,1,1},{2,2,2},{3,3,3}})

	e2a := r2.GetValue(0,0)
	if e2a != 1 {
		t.Error("Example 2a is not correct.")	
	}

	r2.UpdateSubrectangle(0, 0, 2, 2, 100)
	
	e2b := r2.GetValue(0,0)
	if e2b != 100 {
		t.Error("Example 2b is not correct.")
	}

	e2c := r2.GetValue(2,2)
	if e2c != 100 {
		t.Error("Example 2c is not correct.")
	}

	r2.UpdateSubrectangle(1, 1, 2, 2, 20)
	
	e2d := r2.GetValue(2,2)
	if e2d != 20 {
		t.Error("Example 2d is not correct.")
	}
}
