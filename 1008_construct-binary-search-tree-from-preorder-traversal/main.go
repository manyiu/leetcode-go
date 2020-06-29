package main

// TreeNode is the Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}

	t := &TreeNode{Val: preorder[0]}

	for i := 1; i < len(preorder); i++ {
		n := t
		for {
			if preorder[i] > n.Val {
				if n.Right == nil {
					n.Right = &TreeNode{Val: preorder[i]}
					break
				} else {
					n = n.Right
				}
			} else {
				if n.Left == nil {
					n.Left = &TreeNode{Val: preorder[i]}
					break
				} else {
					n = n.Left
				}
			}
		}
	}

	return t
}

func main() {
	bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
}
