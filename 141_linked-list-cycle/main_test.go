package linkedlistcycle

import "testing"

func TestHasCycle(t *testing.T) {
	t.Run("should return false when the list is empty", func(t *testing.T) {
		var head *ListNode

		got := hasCycle(head)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should return false when the list has only one node", func(t *testing.T) {
		head := &ListNode{Val: 1}

		got := hasCycle(head)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should return false when the list has no cycle", func(t *testing.T) {
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next = &ListNode{Val: 4}

		got := hasCycle(head)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should return true when the list has a cycle", func(t *testing.T) {
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next = head.Next

		got := hasCycle(head)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
