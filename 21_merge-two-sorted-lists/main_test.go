package mergetwosortedlists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l3 := mergeTwoLists(l1, l2)

	if l3.Val != 1 || l3.Next.Val != 1 || l3.Next.Next.Val != 2 || l3.Next.Next.Next.Val != 3 || l3.Next.Next.Next.Next.Val != 4 || l3.Next.Next.Next.Next.Next.Val != 4 {
		t.Error("Expected 1 1 2 3 4 4, got", l3.Val, l3.Next.Val, l3.Next.Next.Val, l3.Next.Next.Next.Val, l3.Next.Next.Next.Next.Val, l3.Next.Next.Next.Next.Next.Val)
	}
}
