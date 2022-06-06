package mergetwosortedlists

// ListNode Implement
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l := &ListNode{}
	h := l

	for list1 != nil || list2 != nil {
		if list1 == nil {
			h.Next = list2
			break
		}

		if list2 == nil {
			h.Next = list1
			break
		}

		if list1.Val >= list2.Val {
			h.Next = list2
			list2 = list2.Next
		} else {
			h.Next = list1
			list1 = list1.Next
		}

		h = h.Next
	}

	return l.Next
}
