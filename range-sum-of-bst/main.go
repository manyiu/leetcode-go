package main

import "fmt"

// TreeNode is the Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	l := []*TreeNode{root}

	if l[0] == nil {
		return 0
	}

	sum := 0

	for len(l) > 0 {
		if l[0].Left != nil {
			l = append(l, l[0].Left)
		}
		if l[0].Right != nil {
			l = append(l, l[0].Right)
		}
		if l[0].Val >= L && l[0].Val <= R {
			sum += l[0].Val
		}
		l = l[1:]
	}

	return sum
}

func main() {
	e1 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	e2 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Println(rangeSumBST(e1, 7, 15))
	fmt.Println(rangeSumBST(e2, 6, 10))
}
