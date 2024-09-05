package implementqueueusingstacks

import "testing"

func TestMyQueue(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	if p := obj.Pop(); p != 1 {
		t.Errorf("expect 1, got %d", p)
	}
	obj.Push(5)
	if p := obj.Peek(); p != 2 {
		t.Errorf("expect 2, got %d", p)
	}
}
