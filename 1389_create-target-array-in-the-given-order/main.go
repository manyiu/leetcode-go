package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	r := make([]int, 0, len(nums))

	for i, v := range nums {
		t := append([]int{v}, r[index[i]:]...)
		r = append(r[:index[i]], t...)
	}

	return r
}

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
	fmt.Println(createTargetArray([]int{1}, []int{0}))
}
