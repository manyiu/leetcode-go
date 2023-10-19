package minstack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if minStack.GetMin() != -3 {
		t.Fatalf("GetMin() should be -3, but get %d", minStack.GetMin())
	}
	minStack.Pop()
	if minStack.Top() != 0 {
		t.Fatalf("Top() should be 0, but get %d", minStack.Top())
	}
	if minStack.GetMin() != -2 {
		t.Fatalf("GetMin() should be -2, but get %d", minStack.GetMin())
	}
}
