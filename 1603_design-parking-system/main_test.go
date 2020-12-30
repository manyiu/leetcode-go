package main

import "testing"

func TestParkingSystem(t *testing.T) {
	p := Constructor(1, 1, 0)

	e1 := p.AddCar(1)
	if e1 != true {
		t.Error("AddCar #1 not correct")
	}

	e2 := p.AddCar(2)
	if e2 != true {
		t.Error("AddCar #2 not correct")
	}

	e3 := p.AddCar(3)
	if e3 != false {
		t.Error("AddCar #3 not correct")
	}

	e4 := p.AddCar(1)
	if e4 != false {
		t.Error("AddCar #4 not correct")
	}
}