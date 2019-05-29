package main

import "fmt"

// TreeNode is the definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

// func insertIntoBST(root *TreeNode, val int) *TreeNode {
// 	h := root
// 	n := TreeNode{
// 		Val: val,
// 	}

// 	for {
// 		if val > h.Val {
// 			if h.Right == nil {
// 				h.Right = &n
// 				break
// 			} else {
// 				h = h.Right
// 			}
// 		} else {
// 			if h.Left == nil {
// 				h.Left = &n
// 				break
// 			} else {
// 				h = h.Left
// 			}
// 		}
// 	}

// 	return root
// }

func main() {
	n5 := TreeNode{
		Val: 3,
	}

	n4 := TreeNode{
		Val: 1,
	}

	n3 := TreeNode{
		Val: 7,
	}

	n2 := TreeNode{
		Val:   2,
		Left:  &n4,
		Right: &n5,
	}

	n1 := TreeNode{
		Val:   4,
		Left:  &n2,
		Right: &n3,
	}

	fmt.Println(insertIntoBST(&n1, 5))
}
