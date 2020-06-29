package main

import "fmt"

func subtractProductAndSum(n int) int {
	p := 1
	s := 0

	for n > 0 {
		m := n % 10
		p *= m
		s += m

		n = (n - m) / 10
	}

	return p - s
}

func main() {
	fmt.Println(subtractProductAndSum(234))
	fmt.Println(subtractProductAndSum(4421))
}
