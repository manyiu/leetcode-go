package main

import "fmt"

func findNumbers(nums []int) int {
	c := 0

	for _, v := range nums {
		d := 0

		for v > 0 {
			v /= 10
			d++
		}

		if d%2 == 0 {
			c++
		}
	}

	return c
}

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}
