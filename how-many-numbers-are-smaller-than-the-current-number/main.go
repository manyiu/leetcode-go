package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	r := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		c := 0

		for _, v := range nums {
			if v < nums[i] {
				c++
			}
		}

		r[i] = c
	}

	return r
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}
