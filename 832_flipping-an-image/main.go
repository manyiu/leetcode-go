package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		for i := 0; i < len(a)/2; i++ {
			a[i], a[len(a)-1-i] = a[len(a)-1-i]^1, a[i]^1
		}
		if len(a)%2 == 1 {
			a[len(a)/2] ^= 1
		}
	}

	return A
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}
