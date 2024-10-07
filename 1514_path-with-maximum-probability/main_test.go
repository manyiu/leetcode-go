package pathwithmaximumprobability

import "testing"

func TestMaxProbability(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		edges     [][]int
		succProb  []float64
		startNode int
		endNode   int
		expect    float64
	}{
		{
			name:      "test 1",
			n:         3,
			edges:     [][]int{{0, 1}, {1, 2}, {0, 2}},
			succProb:  []float64{0.5, 0.5, 0.2},
			startNode: 0,
			endNode:   2,
			expect:    0.25,
		},
		{
			name:      "test 2",
			n:         3,
			edges:     [][]int{{0, 1}, {1, 2}, {0, 2}},
			succProb:  []float64{0.5, 0.5, 0.3},
			startNode: 0,
			endNode:   2,
			expect:    0.3,
		},
		{
			name:      "test 3",
			n:         3,
			edges:     [][]int{{0, 1}},
			succProb:  []float64{0.5},
			startNode: 0,
			endNode:   2,
			expect:    0,
		},
		{
			name:      "test 4",
			n:         3,
			edges:     [][]int{{0, 1}},
			succProb:  []float64{0.5},
			startNode: 0,
			endNode:   1,
			expect:    0.5,
		},
		{
			name:      "test 5",
			n:         5,
			edges:     [][]int{{1, 4}, {2, 4}, {0, 4}, {0, 3}, {0, 2}, {2, 3}},
			succProb:  []float64{0.37, 0.17, 0.93, 0.23, 0.39, 0.04},
			startNode: 3,
			endNode:   4,
			expect:    0.2139,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxProbability(test.n, test.edges, test.succProb, test.startNode, test.endNode)
			if got != test.expect {
				t.Errorf("expect %f, got %f", test.expect, got)
			}
		})
	}
}
