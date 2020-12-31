package main

// TreeNode struct
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func traversal(n *TreeNode, m map[int]int, l int) {
	if n == nil {
		return
	}

	if _, ok := m[l]; ok {
		m[l] += n.Val
	} else {
		m[l] = n.Val
	}

	if n.Left != nil {
		traversal(n.Left, m, l + 1)
	}

	if n.Right != nil {
		traversal(n.Right, m, l + 1)
	}
}

func deepestLeavesSum(root *TreeNode) int {
    c := make(map[int]int)

	traversal(root, c, 0)

	d := 0

	for i := range c {
		if i > d {
			d = i
		}
	}

	return c[d]
}

func main() {}
