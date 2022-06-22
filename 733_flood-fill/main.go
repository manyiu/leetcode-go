package floodfill

func walk(image [][]int, sr int, sc int, color int, newColor int) {
	if image[sr][sc] == color {
		image[sr][sc] = newColor

		if sr >= 1 {
			walk(image, sr-1, sc, color, newColor)
		}

		if sc >= 1 {
			walk(image, sr, sc-1, color, newColor)
		}

		if sr+1 < len(image) {
			walk(image, sr+1, sc, color, newColor)
		}

		if sc+1 < len(image[sr]) {
			walk(image, sr, sc+1, color, newColor)
		}
	}
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	c := image[sr][sc]

	if c != color {
		walk(image, sr, sc, c, color)
	}

	return image
}
