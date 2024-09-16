package longestcommonsubsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"

	expected := 3
	result := longestCommonSubsequence(text1, text2)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	text1 = "abc"
	text2 = "abc"

	expected = 3
	result = longestCommonSubsequence(text1, text2)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	text1 = "abc"
	text2 = "def"

	expected = 0
	result = longestCommonSubsequence(text1, text2)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	text1 = "psnw"
	text2 = "vozsh"

	expected = 1
	result = longestCommonSubsequence(text1, text2)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

}
