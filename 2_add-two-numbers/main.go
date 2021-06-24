package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	p := h
	s := 0

	for l1 != nil || l2 != nil || s > 0 {
		switch {
		case l1 != nil && l2 != nil:
			p.Val = l1.Val + l2.Val + s
			l1 = l1.Next
			l2 = l2.Next
		case l1 != nil && l2 == nil:
			p.Val = l1.Val + s
			l1 = l1.Next
		case l1 == nil && l2 != nil:
			p.Val = l2.Val + s
			l2 = l2.Next
		default:
			p.Val = s
		}

		if p.Val >= 10 {
			p.Val -= 10
			s = 1
		} else {
			s = 0
		}

		if l1 != nil || l2 != nil || s > 0 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}

	return h
}

func main() {
	e1l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	e1l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	addTwoNumbers(e1l1, e1l2)
}
