package binarytreelevelordertraversal

import "testing"

func TestLevelOrder(t *testing.T) {
	t.Run("should return empty slice when root is nil", func(t *testing.T) {
		got := levelOrder(nil)
		want := [][]int{}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should return slice of slice of int", func(t *testing.T) {
		root := &TreeNode{Val: 3}
		root.Left = &TreeNode{Val: 9}
		root.Right = &TreeNode{Val: 20}
		root.Right.Left = &TreeNode{Val: 15}
		root.Right.Right = &TreeNode{Val: 7}

		got := levelOrder(root)
		want := [][]int{{3}, {9, 20}, {15, 7}}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
