package main

import "testing"

func TestDeepestLeavesSum(t *testing.T) {
	ns := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	e := deepestLeavesSum(ns)

	if e != 15 {
		t.Error("Example is not correct")
	}
}