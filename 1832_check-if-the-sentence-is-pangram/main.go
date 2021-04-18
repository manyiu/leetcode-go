package main

func checkIfPangram(sentence string) bool {
	s := 0
	m := map[rune]bool{}

	for _, v := range sentence {
		if _, ok := m[v]; !ok {
			m[v] = true
			s++
		}
	}

	if s == 26 {
		return true
	} else {
		return false
	}
}

func main() {}
