package subsets

import "testing"

func TestSubsets(t *testing.T) {
	t.Run("should return all subsets", func(t *testing.T) {
		got := subsets([]int{1, 2, 3})
		want := [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}

		if len(got) != len(want) {
			t.Fatalf("got %v, want %v", got, want)
		}

		for i := range got {
			if len(got[i]) != len(want[i]) {
				t.Fatalf("got %v, want %v", got, want)
			}

			for j := range got[i] {
				if got[i][j] != want[i][j] {
					t.Fatalf("got %v, want %v", got, want)
				}
			}
		}
	})
}
