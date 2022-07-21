package balancedBinaryTree

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getTreeNodeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getTreeNodeHeight(root.Left), getTreeNodeHeight(root.Right))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := getTreeNodeHeight(root.Left)
	rightHeight := getTreeNodeHeight(root.Right)

	if absDiff(leftHeight, rightHeight) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}
