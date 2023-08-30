package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDepth int

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left)
	right := depth(root.Right)

	maxDepth = max(maxDepth, left+right)

	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDepth = 0

	depth(root)

	return maxDepth
}
