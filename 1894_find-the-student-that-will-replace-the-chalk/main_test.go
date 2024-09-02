package findthestudentthatwillreplacethechalk

import "testing"

func TestChalkReplacer(t *testing.T) {
	testCases := []struct {
		chalk []int
		k     int
		want  int
	}{
		{chalk: []int{5, 1, 5}, k: 22, want: 0},
		{chalk: []int{3, 4, 1, 2}, k: 25, want: 1},
		{chalk: []int{1}, k: 1000000000, want: 0},
	}

	for _, tc := range testCases {
		if got := chalkReplacer(tc.chalk, tc.k); got != tc.want {
			t.Errorf("for chalk=%v and k=%d, got %d, want %d", tc.chalk, tc.k, got, tc.want)
		}
	}
}
