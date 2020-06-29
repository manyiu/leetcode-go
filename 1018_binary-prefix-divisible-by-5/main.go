package main

import (
	"fmt"
)

func prefixesDivBy5(A []int) []bool {
	r := make([]bool, len(A))
	a := 0

	for i, v := range A {
		a = a<<1 + v
		r[i] = (a%5 == 0)
		a %= 10
	}

	return r
}

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{0, 1, 1, 1, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1, 0, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}))
}
