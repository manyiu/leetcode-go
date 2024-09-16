package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "example 1",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "example 2",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "example 3",
			m:    7,
			n:    3,
			want: 28,
		},
		{
			name: "example 4",
			m:    3,
			n:    3,
			want: 6,
		},
		{
			name: "example 5",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "example 6",
			m:    1,
			n:    2,
			want: 1,
		},
		{
			name: "example 7",
			m:    2,
			n:    1,
			want: 1,
		},
		{
			name: "example 8",
			m:    2,
			n:    2,
			want: 2,
		},
		{
			name: "example 9",
			m:    2,
			n:    3,
			want: 3,
		},
		{
			name: "example 10",
			m:    3,
			n:    1,
			want: 1,
		},
		{
			name: "example 11",
			m:    1,
			n:    3,
			want: 1,
		},
		{
			name: "example 12",
			m:    4,
			n:    4,
			want: 20,
		},
		{
			name: "example 13",
			m:    4,
			n:    5,
			want: 35,
		},
		{
			name: "example 14",
			m:    23,
			n:    12,
			want: 193536720,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths(%v, %v) = %v, want %v", tt.m, tt.n, got, tt.want)
			}
		})
	}
}
