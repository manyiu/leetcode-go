package main

import "fmt"

func repeatedNTimes(A []int) int {
	m := map[int]int{}

	for _, n := range A {
		if _, ok := m[n]; ok {
			return n
		}

		m[n] = 0
	}

	return 0
}

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}
