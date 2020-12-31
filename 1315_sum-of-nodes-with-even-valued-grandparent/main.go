package main

// TreeNode struct
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func traversal(c *int, n *TreeNode, parent *TreeNode, grandparent *TreeNode) {
	if n == nil {
		return
	}
	
	if grandparent != nil && grandparent.Val % 2 == 0 {
		*c += n.Val
	}

	if n.Left != nil {
		traversal(c, n.Left, n, parent)
	}

	if n.Right != nil {
		traversal(c, n.Right, n, parent)
	}
}


func sumEvenGrandparent(root *TreeNode) int {
    s := 0

	traversal(&s, root, nil, nil)

	return s
}

func main() {}
