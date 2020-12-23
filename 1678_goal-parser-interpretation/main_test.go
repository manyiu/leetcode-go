package main

import "testing"

func TestInterpret(t *testing.T) {
	e1 := interpret("G()(al)")
	if e1 != "Goal" {
		t.Error("Example 1 not correct")
	}

	e2 := interpret("G()()()()(al)")
	if e2 != "Gooooal" {
		t.Error("Example 2 not correct")
	}

	e3 := interpret("(al)G(al)()()G")
	if e3 != "alGalooG" {
		t.Error("Example 3 not correct")
	}
}