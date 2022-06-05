package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		output int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		if maxProfit(tt.prices) != tt.output {
			t.Errorf("maxProfit(%v) return %v, want %v", tt.prices, maxProfit(tt.prices), tt.output)
		}
	}
}
