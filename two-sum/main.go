package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		d := target - n
		if _, e := m[d]; e {
			return []int{m[d], i}
		} else {
			m[n] = i
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
