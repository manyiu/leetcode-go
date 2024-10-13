package minimumcostfortickets

import "testing"

func TestMincostTickets(t *testing.T) {
	tests := []struct {
		days   []int
		costs  []int
		output int
	}{
		{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
		{[]int{1, 5, 8, 9, 10, 12, 13, 16, 17, 18, 19, 20, 23, 24, 29}, []int{3, 12, 54}, 39},
		{[]int{1, 2, 3, 4, 6, 8, 9, 10, 13, 14, 16, 17, 19, 21, 24, 26, 27, 28, 29}, []int{3, 14, 50}, 50},
	}

	for _, test := range tests {
		ret := mincostTickets(test.days, test.costs)
		if ret != test.output {
			t.Errorf("Got %d for input days = %v, costs = %v; expected %d", ret, test.days, test.costs, test.output)
		}
	}
}
