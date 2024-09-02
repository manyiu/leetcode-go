package designlinkedlist

import (
	"testing"
)

func TestMyLinkedList(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)

	if obj.Get(4) != 4 {
		t.Fatal("error")
	}

	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)

	if obj.Get(6) != 0 {
		t.Fatal("error")
	}
}
