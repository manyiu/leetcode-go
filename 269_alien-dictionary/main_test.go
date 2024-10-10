package aliendictionary

import "testing"

func TestForeignDictionary(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{
			[]string{"wrt", "wrf", "er", "ett", "rftt"},
			"wertf",
		},
		{
			[]string{"z", "x"},
			"zx",
		},
		{
			[]string{"z", "x", "z"},
			"",
		},
	}

	for _, test := range tests {
		got := foreignDictionary(test.words)
		if got != test.want {
			t.Errorf("foreignDictionary(%v) = %v; want %v\n", test.words, got, test.want)
		}
	}
}
