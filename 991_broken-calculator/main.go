package main

import "fmt"

func brokenCalc(X int, Y int) int {
	c := 0

	for X < Y {
		if Y%2 == 0 {
			Y /= 2
		} else {
			Y++
		}
		c++
	}

	return c + (X - Y)
}

func main() {
	fmt.Println(brokenCalc(2, 3))
	fmt.Println(brokenCalc(5, 8))
	fmt.Println(brokenCalc(3, 10))
	fmt.Println(brokenCalc(1024, 1))
}
