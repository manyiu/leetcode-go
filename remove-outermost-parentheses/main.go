package main

import "fmt"

func removeOuterParentheses(S string) string {
	o := 0
	c := 0
	n := ""

	for _, b := range S {
		r := false

		if o == c {
			r = true
		}

		if b == '(' {
			o++
		} else {
			c++
		}

		if o == c {
			r = true
		}

		if r != true {
			n += string(b)
		}
	}

	return n
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}
