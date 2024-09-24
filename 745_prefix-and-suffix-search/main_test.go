package prefixandsuffixsearch

import "testing"

func TestWordFilter(t *testing.T) {
	words := []string{"apple", "banana", "orange", "pear", "peach", "pineapple", "plum", "pomegranate"}
	wordFilter := Constructor(words)

	if got := wordFilter.F("a", "e"); got != 0 {
		t.Errorf("wordFilter.F() = %v, want %v", got, 0)
	}

	if got := wordFilter.F("p", "e"); got != 7 {
		t.Errorf("wordFilter.F() = %v, want %v", got, 7)
	}

	if got := wordFilter.F("pe", "ch"); got != 4 {
		t.Errorf("wordFilter.F() = %v, want %v", got, 4)
	}
}
