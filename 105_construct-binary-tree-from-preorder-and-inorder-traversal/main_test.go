package constructbinarytreefrompreorderandinordertraversal

import "testing"

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)

	if root.Val != 3 {
		t.Errorf("got %v want %v given", root.Val, 3)
	}

	if root.Left.Val != 9 {
		t.Errorf("got %v want %v given", root.Left.Val, 9)
	}

	if root.Right.Val != 20 {
		t.Errorf("got %v want %v given", root.Right.Val, 20)
	}

	if root.Right.Left.Val != 15 {
		t.Errorf("got %v want %v given", root.Right.Left.Val, 15)
	}

	if root.Right.Right.Val != 7 {
		t.Errorf("got %v want %v given", root.Right.Right.Val, 7)
	}
}
