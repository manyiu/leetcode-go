package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	j := make(map[rune]bool)
	v := 0

	for _, c := range J {
		j[c] = true
	}

	for _, c := range S {
		if _, e := j[c]; e {
			v++
		}
	}

	return v
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
