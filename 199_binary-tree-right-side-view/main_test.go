package binarytreerightsideview

import "testing"

func TestRightSideView(t *testing.T) {
	// Test case 1
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	expected := []int{1, 3, 4}
	got := rightSideView(root)
	if len(got) != len(expected) {
		t.Errorf("Got: %v, want: %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("Got: %v, want: %v", got, expected)
		}
	}

	// Test case 2
	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	expected = []int{1, 2, 3}
	got = rightSideView(root)
	if len(got) != len(expected) {
		t.Errorf("Got: %v, want: %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("Got: %v, want: %v", got, expected)
		}
	}

	// Test case 3
	root = nil
	expected = []int{}
	got = rightSideView(root)
	if len(got) != len(expected) {
		t.Errorf("Got: %v, want: %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("Got: %v, want: %v", got, expected)
		}
	}
}
