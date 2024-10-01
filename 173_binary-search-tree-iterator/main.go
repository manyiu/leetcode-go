package binarysearchtreeiterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root    *TreeNode
	current *TreeNode
	stack   []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		root:    root,
		current: root,
		stack:   []*TreeNode{},
	}
}

func (this *BSTIterator) Next() int {
	for {
		if this.current != nil {
			this.stack = append(this.stack, this.current)
			this.current = this.current.Left

		} else {
			this.current = this.stack[len(this.stack)-1]
			result := this.current.Val
			this.stack = this.stack[:len(this.stack)-1]
			this.current = this.current.Right
			return result
		}
	}
}

func (this *BSTIterator) HasNext() bool {
	return this.current != nil || len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
