package rottingoranges

func orangesRotting(grid [][]int) int {
	freshCount := 0
	queue := [][2]int{}
	minute := 0

	for r, row := range grid {
		for c, v := range row {
			if v == 1 {
				freshCount++
			} else if v == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	rowLength := len(grid)
	columnLength := len(grid[0])

	steps := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 && freshCount > 0 {
		for _, o := range queue {
			for _, step := range steps {
				nr := o[0] + step[0]
				nc := o[1] + step[1]

				if nr >= 0 && nr < rowLength && nc >= 0 && nc < columnLength && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					queue = append(queue, [2]int{nr, nc})
					freshCount--
				}
			}

			queue = queue[1:]
		}

		minute++
	}

	if freshCount > 0 {
		return -1
	}

	return minute
}
