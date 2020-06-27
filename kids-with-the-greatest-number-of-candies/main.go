package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := 0
	r := make([]bool, len(candies))

	for _, v := range candies {
		if v > m {
			m = v
		}
	}

	for i, v := range candies {
		if v+extraCandies >= m {
			r[i] = true
		}
	}

	return r
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}
