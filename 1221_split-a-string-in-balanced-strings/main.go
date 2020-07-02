package main

import "fmt"

func balancedStringSplit(s string) int {
	c := 0
	g := 0

	for _, v := range s {
		if v == 'L' {
			c--
		} else {
			c++
		}

		if c == 0 {
			g++
		}
	}

	return g
}

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(balancedStringSplit("RLLLLRRRLR"))
	fmt.Println(balancedStringSplit("LLLLRRRR"))
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))
}
