package main

// SubrectangleQueries struct
type SubrectangleQueries struct {
    value [][]int
}

// Constructor function
func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle}
}

// UpdateSubrectangle function
func (s *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			s.value[r][c] = newValue
		}
	}
}

// GetValue function
func (s *SubrectangleQueries) GetValue(row int, col int) int {
    return s.value[row][col] 
}

func main() {}
