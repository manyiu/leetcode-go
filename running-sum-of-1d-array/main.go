package main

import "fmt"

func runningSum(nums []int) []int {
	r := make([]int, len(nums))
	a := 0

	for i, v := range nums {
		a += v
		r[i] = a
	}

	return r
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}
