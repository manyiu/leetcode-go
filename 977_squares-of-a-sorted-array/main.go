package main

import "fmt"

func square(s []int) []int {
	for i, v := range s {
		s[i] = v * v
	}

	return s
}

func mergeSort(s []int) []int {
	if len(s) > 1 {
		m := len(s) / 2
		return merge(mergeSort(s[:m]), mergeSort(s[m:]))
	}

	return s
}

func merge(a, b []int) []int {
	s := []int{}

	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			s = append(s, a[0])
			a = a[1:]
		} else {
			s = append(s, b[0])
			b = b[1:]
		}
	}

	for i := 0; i < len(a); i++ {
		s = append(s, a[i])
	}

	for j := 0; j < len(b); j++ {
		s = append(s, b[j])
	}

	return s
}

func sortedSquares(A []int) []int {
	return mergeSort(square(A))
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
