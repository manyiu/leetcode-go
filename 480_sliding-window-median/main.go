package slidingwindowmedian

import (
	"container/heap"
)

type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func medianSlidingWindow(nums []int, k int) []float64 {
	outMap := map[int]int{}

	result := make([]float64, len(nums)-k+1)

	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}

	heap.Init(minHeap)
	heap.Init(maxHeap)

	for i := 0; i < k; i++ {
		heap.Push(maxHeap, nums[i])
	}

	for len(*maxHeap) > len(*minHeap) {
		heap.Push(minHeap, heap.Pop(maxHeap))
	}

	if k%2 == 1 {
		result[0] = float64((*minHeap)[0])
	} else {
		result[0] = (float64((*minHeap)[0]) + float64((*maxHeap)[0])) / 2
	}

	for curr := k; curr < len(nums); curr++ {
		incoming := nums[curr]
		outgoing := nums[curr-k]

		outMap[outgoing]++

		balance := 0

		if float64(outgoing) < result[curr-k] {
			balance--
		} else {
			balance++
		}

		if float64(incoming) < result[curr-k] {
			heap.Push(maxHeap, incoming)
			balance++
		} else {
			heap.Push(minHeap, incoming)
			balance--
		}

		if balance > 0 {
			heap.Push(minHeap, heap.Pop(maxHeap))
		} else if balance < 0 {
			heap.Push(maxHeap, heap.Pop(minHeap))
		}

		for len(*minHeap) > 0 && outMap[(*minHeap)[0]] > 0 {
			outMap[(*minHeap)[0]]--
			heap.Pop(minHeap)
		}

		for len(*maxHeap) > 0 && outMap[(*maxHeap)[0]] > 0 {
			outMap[(*maxHeap)[0]]--
			heap.Pop(maxHeap)
		}

		if k%2 == 1 {
			result[curr-k+1] = float64((*minHeap)[0])
		} else {
			result[curr-k+1] = (float64((*minHeap)[0]) + float64((*maxHeap)[0])) / 2
		}
	}

	return result
}
