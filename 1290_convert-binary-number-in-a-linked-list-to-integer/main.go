package main

import "fmt"

// ListNode definition
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	v := 0

	for ; head != nil; head = head.Next {
		v = v<<1 + head.Val
	}

	return v
}

func main() {
	e1n2 := &ListNode{
		Val: 1,
	}

	e1n1 := &ListNode{
		Val:  0,
		Next: e1n2,
	}

	e1n0 := &ListNode{
		Val:  1,
		Next: e1n1,
	}

	fmt.Println(getDecimalValue(e1n0))

	e2n0 := &ListNode{
		Val: 0,
	}

	fmt.Println(getDecimalValue(e2n0))

	e3n0 := &ListNode{
		Val: 1,
	}

	fmt.Println(getDecimalValue(e3n0))

	e4n14 := &ListNode{
		Val: 0,
	}

	e4n13 := &ListNode{
		Val:  0,
		Next: e4n14,
	}

	e4n12 := &ListNode{
		Val:  0,
		Next: e4n13,
	}

	e4n11 := &ListNode{
		Val:  0,
		Next: e4n12,
	}

	e4n10 := &ListNode{
		Val:  0,
		Next: e4n11,
	}

	e4n9 := &ListNode{
		Val:  0,
		Next: e4n10,
	}

	e4n8 := &ListNode{
		Val:  1,
		Next: e4n9,
	}

	e4n7 := &ListNode{
		Val:  1,
		Next: e4n8,
	}

	e4n6 := &ListNode{
		Val:  1,
		Next: e4n7,
	}

	e4n5 := &ListNode{
		Val:  0,
		Next: e4n6,
	}

	e4n4 := &ListNode{
		Val:  0,
		Next: e4n5,
	}

	e4n3 := &ListNode{
		Val:  1,
		Next: e4n4,
	}

	e4n2 := &ListNode{
		Val:  0,
		Next: e4n3,
	}

	e4n1 := &ListNode{
		Val:  0,
		Next: e4n2,
	}

	e4n0 := &ListNode{
		Val:  1,
		Next: e4n1,
	}

	fmt.Println(getDecimalValue(e4n0))

	e5n1 := &ListNode{
		Val: 0,
	}

	e5n0 := &ListNode{
		Val:  0,
		Next: e5n1,
	}

	fmt.Println(getDecimalValue(e5n0))
}
