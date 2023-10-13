package evaluatereversepolishnotation

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}

	for _, tt := range tests {
		got := evalRPN(tt.tokens)

		if got != tt.want {
			t.Errorf("got: %v, want: %v, tokens: %v", got, tt.want, tt.tokens)
		}
	}
}
