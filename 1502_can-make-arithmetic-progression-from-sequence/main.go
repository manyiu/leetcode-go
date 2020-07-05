package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	d := 0

	if len(arr) >= 2 {
		d = arr[0] - arr[1]
	}

	for i := 1; i < len(arr)-1; i++ {
		if arr[i]-arr[i+1] != d {
			return false
		}
	}

	return true
}

func main() {}
