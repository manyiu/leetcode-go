package binarytreepostordertraversal

import "testing"

func TestPostorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	got := postorderTraversal(root)
	want := []int{3, 2, 1}
	if len(got) != len(want) {
		t.Fatalf("got len(got) = %v, want %v", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Right.Left = &TreeNode{Val: 6}
	root.Left.Right.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 8}
	root.Right.Right.Left = &TreeNode{Val: 9}

	got = postorderTraversal(root)
	want = []int{4, 6, 7, 5, 2, 9, 8, 3, 1}
	if len(got) != len(want) {
		t.Fatalf("got len(got) = %v, want %v", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
