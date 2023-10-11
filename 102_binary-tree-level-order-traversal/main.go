package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	working := []*TreeNode{root}

	for len(working) > 0 {
		next := []*TreeNode{}
		current := []int{}

		for _, node := range working {
			current = append(current, node.Val)

			if node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		result = append(result, current)
		working = next
	}

	return result
}
