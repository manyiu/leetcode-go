package main

import "testing"

func TestMinPartitions(t *testing.T) {
	e1 := minPartitions("32")
	if e1 != 3 {
		t.Error("Example 1 is not correct.")
	}

	e2 := minPartitions("82734")
	if e2 != 8 {
		t.Error("Example 1 is not correct.")
	}

	e3 := minPartitions("27346209830709182346")
	if e3 != 9 {
		t.Error("Example 1 is not correct.")
	}
}
