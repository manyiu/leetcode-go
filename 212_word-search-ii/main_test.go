package wordsearchii

import "testing"

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}

func TestFindWords(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	expected := []string{"oath", "eat"}
	if got := findWords(board, words); !equal(got, expected) {
		t.Errorf("findWords() = %v, want %v", got, expected)
	}

	board = [][]byte{
		{'o', 'a', 'b', 'n'},
		{'o', 't', 'a', 'e'},
		{'a', 'h', 'k', 'r'},
		{'a', 'f', 'l', 'v'},
	}
	words = []string{"oa", "oaa"}
	expected = []string{"oa", "oaa"}
	if got := findWords(board, words); !equal(got, expected) {
		t.Errorf("findWords() = %v, want %v", got, expected)
	}
}
