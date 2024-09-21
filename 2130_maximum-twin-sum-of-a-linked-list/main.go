package maximumtwinsumofalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow, fast := head, head.Next
	hash := map[int]int{}

	i := 0

	for fast != nil && fast.Next != nil {
		hash[i] = slow.Val
		slow = slow.Next
		fast = fast.Next.Next
		i++
	}

	hash[i] = slow.Val

	max := 0

	for slow != nil && slow.Next != nil {
		slow = slow.Next
		sum := hash[i] + slow.Val

		if sum > max {
			max = sum
		}

		i--
	}

	return max
}
