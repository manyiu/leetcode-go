package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var output []int

	if root == nil {
		return output
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	return append(left, append([]int{root.Val}, right...)...)
}
