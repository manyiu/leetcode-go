package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func indexOf(list []int, val int) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}

	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}

	splitIndex := indexOf(inorder, preorder[0])

	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:splitIndex+1], inorder[:splitIndex]),
		Right: buildTree(preorder[splitIndex+1:], inorder[splitIndex+1:]),
	}

	return root
}
