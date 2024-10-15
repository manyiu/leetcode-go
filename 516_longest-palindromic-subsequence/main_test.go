package longestpalindromicsubsequence

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"empty string", "", 0},
		{"single character", "a", 1},
		{"testcase1", "bbbab", 4},
		{"testcase2", "cbbd", 2},
		{"long string", "euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew", 159},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := longestPalindromeSubseq(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s: %s wanted %d but got %d", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
