package main

func getConcatenation(nums []int) []int {
	r := make([]int, len(nums)*2)

	for i, v := range nums {
		r[i], r[len(nums)+i] = v, v
	}

	return r
}

func main() {}
