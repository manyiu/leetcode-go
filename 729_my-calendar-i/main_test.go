package mycalendari

import "testing"

func TestMyCalendar(t *testing.T) {
	obj := Constructor()
	if !obj.Book(10, 20) {
		t.Error("error")
	}
	if obj.Book(15, 25) {
		t.Error("error")
	}
	if !obj.Book(20, 30) {
		t.Error("error")
	}
}
