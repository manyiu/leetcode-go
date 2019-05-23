package main

import "fmt"

var roman = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	l := len(s)
	v := 0

	for i := 0; i < l; i++ {
		if i+1 < l && roman[rune(s[i])] < roman[rune(s[i+1])] {
			v += roman[rune(s[i+1])] - roman[rune(s[i])]
			i++
		} else {
			v += roman[rune(s[i])]
		}
	}

	return v
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
