package main

import "fmt"

// ListNode Implement
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	l := r
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l.Next = l2
			break
		}
		if l2 == nil {
			l.Next = l1
			break
		}
		if l1.Val > l2.Val {
			l.Next = l2
			l2 = l2.Next
		} else {
			l.Next = l1
			l1 = l1.Next
		}
		l = l.Next
	}

	return r.Next
}

func main() {
	l1c := ListNode{
		Val: 4,
	}

	l1b := ListNode{
		Val:  2,
		Next: &l1c,
	}

	l1a := ListNode{
		Val:  1,
		Next: &l1b,
	}

	l2c := ListNode{
		Val: 4,
	}

	l2b := ListNode{
		Val:  3,
		Next: &l2c,
	}

	l2a := ListNode{
		Val:  1,
		Next: &l2b,
	}

	fmt.Println(mergeTwoLists(&l1a, &l2a))
	fmt.Println(mergeTwoLists(&ListNode{Val: 1}, nil))
}
