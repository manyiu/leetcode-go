package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	merge := lists[0]

	for i := 1; i < len(lists); i++ {
		stage := &ListNode{}
		head := stage

		for merge != nil || lists[i] != nil {
			if merge == nil {
				stage.Next = lists[i]
				lists[i] = lists[i].Next
			} else if lists[i] == nil {
				stage.Next = merge
				merge = merge.Next
			} else if merge.Val < lists[i].Val {
				stage.Next = merge
				merge = merge.Next
			} else {
				stage.Next = lists[i]
				lists[i] = lists[i].Next
			}

			stage = stage.Next
		}

		merge = head.Next
	}

	return merge
}
