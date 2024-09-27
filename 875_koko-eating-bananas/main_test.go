package kokoeatingbananas

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		piles  []int
		h      int
		expect int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
		{[]int{312884470}, 312884469, 2},
		{[]int{1, 1, 1, 999999999}, 10, 142857143},
	}

	for i, test := range tests {
		if minEatingSpeed(test.piles, test.h) != test.expect {
			t.Fatal("test ", i, ": want", test.expect, "but got", minEatingSpeed(test.piles, test.h))
		}
	}
}
