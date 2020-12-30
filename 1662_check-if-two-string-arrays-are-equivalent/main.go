package main

import "strings"

func combineStrings(ss []string) string {
	var sb strings.Builder

	for i := 0; i < len(ss); i++ {
		sb.WriteString(ss[i])
	}

	return sb.String()
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    return combineStrings(word1) == combineStrings(word2)
}

func main() {}
