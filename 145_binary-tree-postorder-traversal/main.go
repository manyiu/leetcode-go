package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	stackNode := []*TreeNode{root}
	stackVisit := []bool{false}

	for len(stackNode) > 0 {
		curr := stackNode[len(stackNode)-1]
		stackNode = stackNode[:len(stackNode)-1]
		visited := stackVisit[len(stackVisit)-1]
		stackVisit = stackVisit[:len(stackVisit)-1]

		if curr != nil {
			if visited {
				result = append(result, curr.Val)
			} else {
				stackNode = append(stackNode, curr)
				stackVisit = append(stackVisit, true)
				stackNode = append(stackNode, curr.Right)
				stackVisit = append(stackVisit, false)
				stackNode = append(stackNode, curr.Left)
				stackVisit = append(stackVisit, false)
			}
		}
	}

	return result
}
