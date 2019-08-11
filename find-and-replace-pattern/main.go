package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	o := []string{}

	for _, word := range words {
		m1 := map[rune]rune{}
		m2 := map[rune]rune{}
		p := true

		for i, char := range word {
			if _, ok1 := m1[char]; ok1 {
				if m1[char] != rune(pattern[i]) || m2[rune(pattern[i])] != char {
					p = false
					break
				}
			} else {
				if _, ok2 := m2[rune(pattern[i])]; ok2 {
					if m2[rune(pattern[i])] != char {
						p = false
						break
					}
				} else {
					m1[char] = rune(pattern[i])
					m2[rune(pattern[i])] = char
				}
			}
		}

		if p == true {
			o = append(o, word)
		}
	}

	return o
}

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}
