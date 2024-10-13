package coinchangeii

import "testing"

func TestChange(t *testing.T) {
	testCases := []struct {
		amount int
		coins  []int
		want   int
	}{
		{amount: 5, coins: []int{1, 2, 5}, want: 4},
		{amount: 3, coins: []int{2}, want: 0},
		{amount: 10, coins: []int{10}, want: 1},
	}

	for _, tc := range testCases {
		if got := change(tc.amount, tc.coins); got != tc.want {
			t.Errorf("amount: %v, coins: %v, got: %v, want: %v", tc.amount, tc.coins, got, tc.want)
		}
	}
}
