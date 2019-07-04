package main

import "fmt"

func sortArrayByParity(A []int) []int {
	n := []int{}

	for _, i := range A {
		if i%2 == 0 {
			n = append([]int{i}, n...)
		} else {
			n = append(n, i)
		}
	}

	return n
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
