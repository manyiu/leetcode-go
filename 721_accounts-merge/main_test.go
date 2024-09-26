package accountsmerge

import "testing"

func equal(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestAccountsMerge(t *testing.T) {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	expected := [][]string{
		{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	result := accountsMerge(accounts)
	if !equal(result, expected) {
		t.Errorf("expected %v, but got %v", expected, result)
	}
}
