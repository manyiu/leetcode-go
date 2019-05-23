package main

import "fmt"

func isValid(s string) bool {
	pair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	valid := true
	closes := []rune{}

	for _, c := range s {
		if b, open := pair[c]; open {
			closes = append([]rune{b}, closes...)
		} else if len(closes) == 0 {
			valid = false
			break
		} else if c == closes[0] {
			closes = closes[1:]
		} else {
			valid = false
			break
		}
	}

	if len(closes) != 0 {
		valid = false
	}

	return valid
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
