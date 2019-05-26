package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x > int(math.Pow(2, 31)-1) || x < int(-math.Pow(2, 31)) {
		return 0
	}

	y := x
	n := 0

	if x < 0 {
		y = -y
	}

	for y > 0 {
		r := y % 10
		n *= 10
		n += r
		y /= 10
	}

	if x < 0 {
		n = -n
	}

	if n > int(math.Pow(2, 31)-1) || n < int(-math.Pow(2, 31)) {
		return 0
	}

	return n
}

func main() {
	fmt.Println(reverse(-1230343000))
}
