package trimabinarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkSameBST(t *testing.T, expected *TreeNode, actual *TreeNode) bool {
	if expected == nil && actual == nil {
		return true
	}

	if expected == nil || actual == nil {
		return false
	}

	return assert.Equal(t, expected.Val, actual.Val) && checkSameBST(t, expected.Left, actual.Left) && checkSameBST(t, expected.Right, actual.Right)
}

func TestTrimBST(t *testing.T) {
	e1input := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	e1expected := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	assert.True(t, checkSameBST(t, e1expected, trimBST(e1input, 1, 2)))

	e2input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	e2expected := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}

	assert.True(t, checkSameBST(t, e2expected, trimBST(e2input, 1, 3)))
}
