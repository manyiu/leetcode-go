package binarytreeinordertraversal

import "testing"

func TestInorderTraversal(t *testing.T) {
	var root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	var got = inorderTraversal(root)
	var want = []int{1, 3, 2}

	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Fatalf("got: %v, want: %v", got, want)
		}
	}
}
