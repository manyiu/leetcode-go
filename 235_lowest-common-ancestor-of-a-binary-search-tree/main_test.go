package lowestCommonAncestorBinarySearchTree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 4}
	root.Right.Right.Left = &TreeNode{Val: 9}
	root.Right.Right.Right = &TreeNode{Val: 11}

	if lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 1}).Val != 3 {
		t.Error("lowestCommonAncestor error")
	}
}
