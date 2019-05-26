package main

import (
	"sort"
)

// TreeNode implement
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	l := []*TreeNode{}
	s := []int{}
	a := []int{}
	m := map[int]int{}

	for len(q) > 0 {
		if q[0].Left != nil {
			q = append(q, q[0].Left)
		}
		if q[0].Right != nil {
			q = append(q, q[0].Right)
		}
		l = append(l, q[0])
		q = q[1:]
	}

	for _, v := range l {
		s = append(s, v.Val)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	for i, v := range s {
		if i == 0 {
			a = append(a, v)
			m[v] = v
		} else {
			a = append(a, v+a[i-1])
			m[v] = v + a[i-1]
		}
	}

	for _, p := range l {
		p.Val = m[p.Val]
	}

	return root
}

func main() {
	t9 := TreeNode{
		Val: 8,
	}

	t8 := TreeNode{
		Val: 3,
	}

	t7 := TreeNode{
		Val:   7,
		Right: &t9,
	}

	t6 := TreeNode{
		Val: 5,
	}

	t5 := TreeNode{
		Val:   2,
		Right: &t8,
	}

	t4 := TreeNode{
		Val: 0,
	}

	t3 := TreeNode{
		Val:   6,
		Left:  &t6,
		Right: &t7,
	}

	t2 := TreeNode{
		Val:   1,
		Left:  &t4,
		Right: &t5,
	}

	t1 := TreeNode{
		Val:   4,
		Left:  &t2,
		Right: &t3,
	}

	bstToGst(&t1)
}
