package longestturbulentsubarray

import "testing"

func TestMaxTurbulenceSize(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "example 1",
			arr:  []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			want: 5,
		},
		{
			name: "example 2",
			arr:  []int{4, 8, 12, 16},
			want: 2,
		},
		{
			name: "example 3",
			arr:  []int{100},
			want: 1,
		},
		{
			name: "example 4",
			arr:  []int{9, 9},
			want: 1,
		},
		{
			name: "example 5",
			arr:  []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24},
			want: 8,
		},
		{
			name: "example 6",
			arr:  []int{37, 199, 60, 296, 257, 248, 115, 31, 273, 176},
			want: 5,
		},
		{
			name: "example 7",
			arr:  []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			want: 1,
		},
		{
			name: "example 8",
			arr:  []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			want: 1,
		},
		{
			name: "example 9",
			arr:  []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTurbulenceSize(tt.arr)
			if got != tt.want {
				t.Errorf("name %s got %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}
