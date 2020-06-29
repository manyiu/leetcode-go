package main

import "fmt"

func decompressRLElist(nums []int) []int {
	r := []int{}

	for i := 0; i < len(nums)-1; i = i + 2 {
		for j := 0; j < nums[i]; j++ {
			r = append(r, nums[i+1])
		}
	}

	return r
}

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}
