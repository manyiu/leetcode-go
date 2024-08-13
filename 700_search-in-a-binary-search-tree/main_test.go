package searchinabinarysearchtree

import "testing"

func TestSearchBST(t *testing.T) {
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
			val:  2,
			want: tree1,
		},
		{
			root: &TreeNode{
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
				},
			},
			val:  5,
			want: nil,
		},
	}

	for _, tt := range tests {
		got := searchBST(tt.root, tt.val)

		if got != tt.want {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}
}
