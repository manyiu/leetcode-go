package countingbits

import "testing"

func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "test 1",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "test 2",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.n)
			if len(got) != len(tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}

			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("got: %v, want: %v", got, tt.want)
				}
			}
		})
	}
}
