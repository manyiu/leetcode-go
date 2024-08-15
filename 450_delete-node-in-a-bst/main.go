package deletenodeinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMinValue(root *TreeNode) int {
	var curr = root

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}

	if key == root.Val {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		var min = findMinValue(root.Right)

		root.Val = min
		root.Right = deleteNode(root.Right, min)
	}

	return root
}
