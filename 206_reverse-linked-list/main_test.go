package reverselinkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "test1",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
		{
			name: "test2",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
		{
			name: "test3",
			head: &ListNode{
				Val: 1,
			},
			want: &ListNode{
				Val: 1,
			},
		},
		{
			name: "test4",
			head: nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
