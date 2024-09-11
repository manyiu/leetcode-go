package maxareaofisland

func count(grid [][]int, r, c int) int {
	if grid[r][c] != 1 {
		return 0
	}

	total := 1

	grid[r][c] = 0

	if r-1 >= 0 {
		total += count(grid, r-1, c)
	}

	if r+1 < len(grid) {
		total += count(grid, r+1, c)
	}

	if c-1 >= 0 {
		total += count(grid, r, c-1)
	}

	if c+1 < len(grid[r]) {
		total += count(grid, r, c+1)
	}

	return total
}

func maxAreaOfIsland(grid [][]int) int {
	max := 0

	for i, row := range grid {
		for j := range row {
			c := count(grid, i, j)

			if c > max {
				max = c
			}
		}
	}

	return max
}
