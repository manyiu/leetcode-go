package main

func countConsistentStrings(allowed string, words []string) int {
	a, c := 0, make([]bool, 26)

	for _, v := range allowed {
		c[v - 'a'] = true
	}

	Word:
		for _, word := range words {
			for _, r := range word {
				if c[r - 'a'] == false {
					continue Word
				}
			}
			
			a++
		}

	return a
}

func main() {}
