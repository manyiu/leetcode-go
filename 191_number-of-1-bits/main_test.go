package numberof1bits

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "test 1",
			n:    0b00000000000000000000000000001011,
			want: 3,
		},
		{
			name: "test 2",
			n:    0b00000000000000000000000010000000,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.n)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
