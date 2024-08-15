package insertintoabinarysearchtree

import "testing"

func TestInsertIntoBST(t *testing.T) {
	tree1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{
			root: &TreeNode{
				Val:  4,
				Left: tree1,
				Right: &TreeNode{
					Val: 7,
				},
			},
			val: 5,
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		got := insertIntoBST(tt.root, tt.val)

		if got.Val != tt.want.Val {
			t.Errorf("got: %v, want: %v", got.Val, tt.want.Val)
		}

		if got.Left.Val != tt.want.Left.Val {
			t.Errorf("got: %v, want: %v", got.Left.Val, tt.want.Left.Val)
		}

		if got.Left.Left.Val != tt.want.Left.Left.Val {
			t.Errorf("got: %v, want: %v", got.Left.Left.Val, tt.want.Left.Left.Val)
		}

		if got.Left.Right.Val != tt.want.Left.Right.Val {
			t.Errorf("got: %v, want: %v", got.Left.Right.Val, tt.want.Left.Right.Val)
		}

		if got.Right.Val != tt.want.Right.Val {
			t.Errorf("got: %v, want: %v", got.Right.Val, tt.want.Right.Val)
		}

		if got.Right.Left.Val != tt.want.Right.Left.Val {
			t.Errorf("got: %v, want: %v", got.Right.Left.Val, tt.want.Right.Left.Val)
		}
	}
}
