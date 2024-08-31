package kthlargestelementinanarray

type minHeap struct {
	k     int
	queue []int
}

func (heap *minHeap) Heapify(nums []int) {
	nums = append([]int{0}, nums...)

	for i := len(nums) / 2; i >= 1; i-- {
		curr := i

		for curr*2 < len(nums) {
			if curr*2+1 < len(nums) && nums[curr] > nums[curr*2+1] && nums[curr*2] > nums[curr*2+1] {
				nums[curr], nums[curr*2+1] = nums[curr*2+1], nums[curr]
				curr = curr*2 + 1
			} else if nums[curr] > nums[curr*2] {
				nums[curr], nums[curr*2] = nums[curr*2], nums[curr]
				curr *= 2
			} else {
				break
			}
		}
	}

	heap.queue = nums
}

func (heap *minHeap) Pop() int {
	popValue := heap.queue[1]
	heap.queue[1] = heap.queue[len(heap.queue)-1]
	heap.queue = heap.queue[:len(heap.queue)-1]

	curr := 1

	for curr*2 < len(heap.queue) {
		if curr*2+1 < len(heap.queue) && heap.queue[curr] > heap.queue[curr*2+1] && heap.queue[curr*2+1] < heap.queue[curr*2] {
			heap.queue[curr*2+1], heap.queue[curr] = heap.queue[curr], heap.queue[curr*2+1]
			curr = curr*2 + 1
		} else if heap.queue[curr] > heap.queue[curr*2] {
			heap.queue[curr], heap.queue[curr*2] = heap.queue[curr*2], heap.queue[curr]
			curr *= 2
		} else {
			break
		}
	}

	return popValue
}

func findKthLargest(nums []int, k int) int {
	mh := &minHeap{
		k:     k,
		queue: []int{},
	}

	mh.Heapify(nums)

	for len(mh.queue)-1 > k {
		mh.Pop()
	}

	return mh.queue[1]
}
