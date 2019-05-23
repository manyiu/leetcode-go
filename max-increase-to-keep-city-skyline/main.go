package main

import "fmt"

func maxIncreaseKeepingSkyline(grid [][]int) int {
	increase := 0
	length := len(grid)
	rMax := make([]int, len(grid))
	cMax := make([]int, len(grid))

	for i := 0; i < length; i++ {
		temp := 0
		for j := 0; j < length; j++ {
			if grid[i][j] > temp {
				temp = grid[i][j]
			}
		}
		rMax[i] = temp

		temp = 0
		for j := 0; j < length; j++ {
			if grid[j][i] > temp {
				temp = grid[j][i]
			}
		}
		cMax[i] = temp
	}

	for i, r := range grid {
		for j, b := range r {
			if b < rMax[i] && b < cMax[j] {
				if rMax[i] < cMax[j] {
					increase += rMax[i] - b
				} else {
					increase += cMax[j] - b
				}
			}
		}
	}

	return increase
}

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
}
