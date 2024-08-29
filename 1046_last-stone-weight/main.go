package laststoneweight

type MaxHeap struct {
	priorityQueue []int
}

func (heap *MaxHeap) Add(val int) int {
	heap.priorityQueue = append(heap.priorityQueue, val)

	curr := len(heap.priorityQueue) - 1

	for curr > 1 {
		if heap.priorityQueue[curr] > heap.priorityQueue[curr/2] {
			heap.priorityQueue[curr], heap.priorityQueue[curr/2] = heap.priorityQueue[curr/2], heap.priorityQueue[curr]
			curr /= 2
		} else {
			break
		}
	}

	if len(heap.priorityQueue) >= 2 {
		return heap.priorityQueue[1]
	}

	return 0
}

func (heap *MaxHeap) Pop() int {
	if len(heap.priorityQueue) <= 1 {
		return 0
	}

	popValue := heap.priorityQueue[1]

	heap.priorityQueue[1] = heap.priorityQueue[len(heap.priorityQueue)-1]
	heap.priorityQueue = append([]int{0}, heap.priorityQueue[1:len(heap.priorityQueue)-1]...)
	curr := 1

	for curr*2 < len(heap.priorityQueue) {
		if curr*2+1 < len(heap.priorityQueue) && heap.priorityQueue[curr*2+1] > heap.priorityQueue[curr*2] && heap.priorityQueue[curr] < heap.priorityQueue[curr*2+1] {
			heap.priorityQueue[curr], heap.priorityQueue[curr*2+1] = heap.priorityQueue[curr*2+1], heap.priorityQueue[curr]
			curr = curr*2 + 1
		} else if heap.priorityQueue[curr] < heap.priorityQueue[curr*2] {
			heap.priorityQueue[curr], heap.priorityQueue[curr*2] = heap.priorityQueue[curr*2], heap.priorityQueue[curr]
			curr *= 2
		} else {
			break
		}
	}

	return popValue
}

func lastStoneWeight(stones []int) int {
	mh := &MaxHeap{
		priorityQueue: []int{0},
	}

	for _, stone := range stones {
		mh.Add(stone)
	}

	for len(mh.priorityQueue) >= 3 {
		a := mh.Pop()
		b := mh.Pop()

		if a > b {
			mh.Add(a - b)
		}
	}

	if len(mh.priorityQueue) == 2 {
		return mh.priorityQueue[1]
	}

	return 0
}
