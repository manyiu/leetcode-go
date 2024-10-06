package networkdelaytime

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name   string
		times  [][]int
		n      int
		k      int
		expect int
	}{
		{
			name:   "test1",
			times:  [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:      4,
			k:      2,
			expect: 2,
		},
		{
			name:   "test2",
			times:  [][]int{{1, 2, 1}},
			n:      2,
			k:      2,
			expect: -1,
		},
		{
			name:   "test3",
			times:  [][]int{{1, 2, 1}},
			n:      2,
			k:      1,
			expect: 1,
		},
		{
			name:   "test4",
			times:  [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}},
			n:      3,
			k:      1,
			expect: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := networkDelayTime(test.times, test.n, test.k)
			if got != test.expect {
				t.Errorf("expect %d, got %d", test.expect, got)
			}
		})
	}
}
