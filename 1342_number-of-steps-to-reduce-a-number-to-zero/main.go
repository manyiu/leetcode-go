package main

import "fmt"

func numberOfSteps(num int) int {
	c := 0

	for {
		if num == 0 {
			return c
		}

		if num%2 != 0 {
			num--
		} else {
			num /= 2
		}

		c++
	}
}

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
}
