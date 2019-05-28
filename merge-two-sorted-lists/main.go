package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := []*ListNode{}
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l = append(l, l2)
			l2 = l2.Next
			break
		}
		if l2 == nil {
			l = append(l, l1)
			l1 = l1.Next
			break
		}
		if l1.Val > l2.Val {
			l = append(l, l2)
			l2 = l2.Next
		} else {
			l = append(l, l1)
			l1 = l1.Next
		}
	}

	for i := 0; i+1 < len(l); i++ {
		l[i].Next = l[i+1]
	}

	if len(l) == 0 {
		return nil
	}

	return l[0]
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
