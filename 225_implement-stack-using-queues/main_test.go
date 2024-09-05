package implementstackusingqueues

import (
	"testing"
)

func TestMyStack(t *testing.T) {
	stack := Constructor()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if got := stack.Top(); got != 3 {
		t.Errorf("Top() = %v; want 3", got)
	}

	if got := stack.Pop(); got != 3 {
		t.Errorf("Pop() = %v; want 3", got)
	}

	if got := stack.Top(); got != 2 {
		t.Errorf("Top() = %v; want 2", got)
	}

	if got := stack.Pop(); got != 2 {
		t.Errorf("Pop() = %v; want 2", got)
	}

	if got := stack.Empty(); got != false {
		t.Errorf("Empty() = %v; want false", got)
	}

	if got := stack.Pop(); got != 1 {
		t.Errorf("Pop() = %v; want 1", got)
	}

	if got := stack.Empty(); got != true {
		t.Errorf("Empty() = %v; want true", got)
	}
}
