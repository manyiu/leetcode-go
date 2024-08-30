package kclosestpointstoorigin

import (
	"math"
)

type minHeap struct {
	k     int
	queue [][]int
}

func (heap *minHeap) Heapify() {
	queue := append([][]int{{0, 0}}, heap.queue...)

	for i := len(queue) / 2; i > 0; i-- {
		curr := i

		for curr*2 < len(queue) {
			if curr*2+1 < len(queue) && math.Pow(float64(queue[curr][0]), 2)+math.Pow(float64(queue[curr][1]), 2) > math.Pow(float64(queue[curr*2+1][0]), 2)+math.Pow(float64(queue[curr*2+1][1]), 2) && math.Pow(float64(queue[curr*2][0]), 2)+math.Pow(float64(queue[curr*2][1]), 2) > math.Pow(float64(queue[curr*2+1][0]), 2)+math.Pow(float64(queue[curr*2+1][1]), 2) {
				queue[curr], queue[curr*2+1] = queue[curr*2+1], queue[curr]
				curr = curr*2 + 1
			} else if math.Pow(float64(queue[curr][0]), 2)+math.Pow(float64(queue[curr][1]), 2) > math.Pow(float64(queue[curr*2][0]), 2)+math.Pow(float64(queue[curr*2][1]), 2) {
				queue[curr], queue[curr*2] = queue[curr*2], queue[curr]
				curr *= 2
			} else {
				break
			}
		}
	}

	heap.queue = queue
}

func (heap *minHeap) Pop() []int {
	if len(heap.queue) <= 1 {
		return []int{0, 0}
	}

	popValue := heap.queue[1]
	heap.queue[1] = heap.queue[len(heap.queue)-1]
	heap.queue = heap.queue[:len(heap.queue)-1]

	curr := 1

	for curr*2 < len(heap.queue) {
		if curr*2+1 < len(heap.queue) && math.Pow(float64(heap.queue[curr][0]), 2)+math.Pow(float64(heap.queue[curr][1]), 2) > math.Pow(float64(heap.queue[curr*2+1][0]), 2)+math.Pow(float64(heap.queue[curr*2+1][1]), 2) && math.Pow(float64(heap.queue[curr*2+1][0]), 2)+math.Pow(float64(heap.queue[curr*2+1][1]), 2) < math.Pow(float64(heap.queue[curr*2][0]), 2)+math.Pow(float64(heap.queue[curr*2][1]), 2) {
			heap.queue[curr], heap.queue[curr*2+1] = heap.queue[curr*2+1], heap.queue[curr]
			curr = curr*2 + 1
		} else if math.Pow(float64(heap.queue[curr][0]), 2)+math.Pow(float64(heap.queue[curr][1]), 2) > math.Pow(float64(heap.queue[curr*2][0]), 2)+math.Pow(float64(heap.queue[curr*2][1]), 2) {
			heap.queue[curr], heap.queue[curr*2] = heap.queue[curr*2], heap.queue[curr]
			curr *= 2
		} else {
			break
		}
	}

	return popValue
}

func kClosest(points [][]int, k int) [][]int {
	heap := &minHeap{
		k:     k,
		queue: points,
	}

	heap.Heapify()

	result := [][]int{}

	for i := 0; i < k; i++ {
		result = append(result, heap.Pop())
	}

	return result
}
