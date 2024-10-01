package binarysearchtreeiterator

import "testing"

func TestBSTIterator(t *testing.T) {
	// root = [7, 3, 15, null, null, 9, 20]
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	obj := Constructor(root)
	param1 := obj.Next()
	if param1 != 3 {
		t.Fatalf("expected 3 but got %d", param1)
	}
	param2 := obj.HasNext()
	if param2 != true {
		t.Fatalf("expected true but got %v", param2)
	}
	param3 := obj.Next()
	if param3 != 7 {
		t.Fatalf("expected 7 but got %d", param3)
	}
	param4 := obj.HasNext()
	if param4 != true {
		t.Fatalf("expected true but got %v", param4)
	}
	param5 := obj.Next()
	if param5 != 9 {
		t.Fatalf("expected 9 but got %d", param5)
	}
	param6 := obj.HasNext()
	if param6 != true {
		t.Fatalf("expected true but got %v", param6)
	}
	param7 := obj.Next()
	if param7 != 15 {
		t.Fatalf("expected 15 but got %d", param7)
	}
	param8 := obj.HasNext()
	if param8 != true {
		t.Fatalf("expected true but got %v", param8)
	}
	param9 := obj.Next()
	if param9 != 20 {
		t.Fatalf("expected 20 but got %d", param9)
	}
	param10 := obj.HasNext()
	if param10 != false {
		t.Fatalf("expected false but got %v", param10)
	}
}
