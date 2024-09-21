package maximumtwinsumofalinkedlist

import "testing"

func TestPairSum(t *testing.T) {
	head := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	got := pairSum(head)
	want := 6
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	head = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	got = pairSum(head)
	want = 7
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 100000,
		},
	}
	got = pairSum(head)
	want = 100001
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
