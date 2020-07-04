package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	c := 0

	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			c++
		}
	}

	return c
}

func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
	fmt.Println(busyStudent([]int{4}, []int{4}, 4))
	fmt.Println(busyStudent([]int{4}, []int{4}, 5))
	fmt.Println(busyStudent([]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7))
	fmt.Println(busyStudent([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5))
}
