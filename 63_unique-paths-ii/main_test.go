package uniquepathsii

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	expected := 2
	result := uniquePathsWithObstacles(obstacleGrid)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

}
