package ipo

import (
	"container/heap"
)

type MinHeap [][2]int
type MaxHeap []int

func (mh MinHeap) Less(i, j int) bool {
	return mh[i][0] < mh[j][0]
}
func (mh MaxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MinHeap) Len() int {
	return len(mh)
}
func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.([2]int))
}
func (mh *MaxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *MinHeap) Pop() interface{} {
	n := len(*mh)
	x := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return x
}
func (mh *MaxHeap) Pop() interface{} {
	n := len(*mh)
	x := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}

	for i := 0; i < len(profits); i++ {
		*minHeap = append(*minHeap, [2]int{capital[i], profits[i]})
	}

	heap.Init(minHeap)
	heap.Init(maxHeap)

	for i := 0; i < k; i++ {
		for len(*minHeap) > 0 && w >= (*minHeap)[0][0] {
			x := heap.Pop(minHeap).([2]int)
			heap.Push(maxHeap, x[1])
		}

		if len(*maxHeap) <= 0 {
			break
		}

		w += heap.Pop(maxHeap).(int)
	}

	return w
}
