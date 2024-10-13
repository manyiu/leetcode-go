package coinchange

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{coins: []int{1, 2, 5}, amount: 11, want: 3},
		{coins: []int{2}, amount: 3, want: -1},
		{coins: []int{1}, amount: 0, want: 0},
		{coins: []int{1}, amount: 1, want: 1},
		{coins: []int{1}, amount: 2, want: 2},
	}

	for _, tc := range testCases {
		if got := coinChange(tc.coins, tc.amount); got != tc.want {
			t.Errorf("coins: %v, amount: %v, got: %v, want: %v", tc.coins, tc.amount, got, tc.want)
		}
	}
}
