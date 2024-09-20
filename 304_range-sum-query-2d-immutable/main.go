package rangesumquery2dimmutable

type NumMatrix struct {
	prefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefixSum := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		row := make([]int, len(matrix[i]))

		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				if j == 0 {
					row[j] = matrix[i][j]
				} else {
					row[j] = row[j-1] + matrix[i][j]
				}
			} else {
				if j == 0 {
					row[j] = prefixSum[i-1][j] + matrix[i][j]
				} else {
					row[j] = prefixSum[i-1][j] + row[j-1] - prefixSum[i-1][j-1] + matrix[i][j]
				}
			}
		}
		prefixSum[i] = row
	}

	return NumMatrix{
		prefixSum: prefixSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	top := 0

	if row1 >= 1 {
		top = this.prefixSum[row1-1][col2]
	}

	left := 0

	if col1 >= 1 {
		left = this.prefixSum[row2][col1-1]
	}

	topLeft := 0

	if row1 >= 1 && col1 >= 1 {
		topLeft = this.prefixSum[row1-1][col1-1]
	}

	return this.prefixSum[row2][col2] - top - left + topLeft
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
