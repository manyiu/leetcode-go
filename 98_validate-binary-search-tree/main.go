package validatebinarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMin(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func findMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func checker(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return checker(root.Left, min, findMin(root.Val, max)) && checker(root.Right, findMax(root.Val, min), max)
}

func isValidBST(root *TreeNode) bool {
	return checker(root, math.MinInt, math.MaxInt)
}
