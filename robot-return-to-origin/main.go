package main

func judgeCircle(moves string) bool {
	m := map[string]int{
		"U": 0,
		"D": 0,
		"L": 0,
		"R": 0,
	}

	for _, d := range moves {
		m[string(d)]++
	}

	if m["U"] == m["D"] && m["L"] == m["R"] {
		return true
	}

	return false
}

func main() {
	judgeCircle("UD")
	judgeCircle("LL")
}
