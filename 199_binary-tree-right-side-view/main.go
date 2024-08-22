package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{root.Val}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		levelStack := []*TreeNode{}

		for _, node := range stack {
			if node.Left != nil {
				levelStack = append(levelStack, node.Left)
			}

			if node.Right != nil {
				levelStack = append(levelStack, node.Right)
			}
		}

		stack = levelStack

		if len(levelStack) > 0 {
			result = append(result, levelStack[len(levelStack)-1].Val)
		}
	}

	return result
}
