package kthsmallestelementinabst

import "testing"

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	got := kthSmallest(root, 1)
	want := 1

	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
