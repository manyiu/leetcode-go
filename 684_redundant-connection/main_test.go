package redundantconnection

import "testing"

func TestFindRedundantConnection(t *testing.T) {
	st := []struct {
		name  string
		edges [][]int
		exp   []int
	}{
		{"testcase1", [][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		{"testcase2", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, []int{1, 4}},
		{"testcase3", [][]int{{16, 25}, {7, 9}, {3, 24}, {10, 20}, {15, 24}, {2, 8}, {19, 21}, {2, 15}, {13, 20}, {5, 21}, {7, 11}, {6, 23}, {7, 16}, {1, 8}, {17, 20}, {4, 19}, {11, 22}, {5, 11}, {1, 16}, {14, 20}, {1, 4}, {22, 23}, {12, 20}, {15, 18}, {12, 16}}, []int{1, 4}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findRedundantConnection(tt.edges)
			if len(out) != len(tt.exp) {
				t.Fatalf("with input edges: %v wanted %v but got %v", tt.edges, tt.exp, out)
			}
			if out[0] != tt.exp[0] || out[1] != tt.exp[1] {
				t.Fatalf("with input edges: %v wanted %v but got %v", tt.edges, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
