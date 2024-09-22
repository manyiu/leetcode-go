package linkedlistcycleii

import "testing"

func TestDetectCycle(t *testing.T) {
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: -4,
				},
			},
		},
	}
	head.Next.Next.Next.Next = head.Next
	got := detectCycle(head)
	want := head.Next
	if got != want {
		t.Errorf("test 1: got %v want %v given", got, want)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	head.Next.Next = head
	got = detectCycle(head)
	want = head
	if got != want {
		t.Errorf("test 2: got %v want %v given", got, want)
	}

	head = &ListNode{
		Val: 1,
	}
	got = detectCycle(head)
	want = nil
	if got != want {
		t.Errorf("test 3: got %v want %v given", got, want)
	}
}
