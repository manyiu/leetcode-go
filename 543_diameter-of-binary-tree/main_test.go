package diameterofbinarytree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	if d := diameterOfBinaryTree(root); d != 3 {
		t.Errorf("diameterOfBinaryTree(%v) = %d; want 3", root, d)
	}
}
