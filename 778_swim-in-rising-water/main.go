package swiminrisingwater

import (
	"container/heap"
)

type MinHeap [][3]int

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i][2] < mh[j][2]
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.([3]int))
}

func (mh *MinHeap) Pop() interface{} {
	n := len(*mh)
	x := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return x
}

func swimInWater(grid [][]int) int {
	added := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		added[i] = make([]bool, len(grid[i]))
	}

	added[0][0] = true

	minHeap := &MinHeap{
		// {row, column, elevation}
		{0, 0, grid[0][0]},
	}

	for len(*minHeap) > 0 {
		popped := heap.Pop(minHeap).([3]int)
		row, column, elevation := popped[0], popped[1], popped[2]

		if row == len(grid)-1 && column == len(grid[row])-1 {
			return elevation
		}

		if row >= 1 && !added[row-1][column] {
			heap.Push(minHeap, [3]int{row - 1, column, max(elevation, grid[row-1][column])})
			added[row-1][column] = true
		}

		if column >= 1 && !added[row][column-1] {
			heap.Push(minHeap, [3]int{row, column - 1, max(elevation, grid[row][column-1])})
			added[row][column-1] = true
		}

		if row < len(grid)-1 && !added[row+1][column] {
			heap.Push(minHeap, [3]int{row + 1, column, max(elevation, grid[row+1][column])})
			added[row+1][column] = true
		}

		if column < len(grid[row])-1 && !added[row][column+1] {
			heap.Push(minHeap, [3]int{row, column + 1, max(elevation, grid[row][column+1])})
			added[row][column+1] = true
		}
	}

	return -1
}
