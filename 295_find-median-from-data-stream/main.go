package findmedianfromdatastream

type MinHeap struct {
	arr []int
}

type MaxHeap struct {
	arr []int
}

func (minHeap *MinHeap) Push(val int) {
	minHeap.arr = append(minHeap.arr, val)

	curr := len(minHeap.arr) - 1

	for curr >= 1 {
		if minHeap.arr[curr] < minHeap.arr[(curr-1)/2] {
			minHeap.arr[curr], minHeap.arr[(curr-1)/2] = minHeap.arr[(curr-1)/2], minHeap.arr[curr]
		} else {
			break
		}

		curr = (curr - 1) / 2
	}
}

func (minHeap *MinHeap) Pop() int {
	result := minHeap.arr[0]

	minHeap.arr[0], minHeap.arr[len(minHeap.arr)-1] = minHeap.arr[len(minHeap.arr)-1], minHeap.arr[0]
	minHeap.arr = minHeap.arr[:len(minHeap.arr)-1]

	curr := 0

	for curr*2+1 < len(minHeap.arr) {
		if curr*2+2 < len(minHeap.arr) && minHeap.arr[curr*2+2] < minHeap.arr[curr*2+1] && minHeap.arr[curr*2+2] < minHeap.arr[curr] {
			minHeap.arr[curr], minHeap.arr[curr*2+2] = minHeap.arr[curr*2+2], minHeap.arr[curr]
			curr = curr*2 + 2
		} else if minHeap.arr[curr] > minHeap.arr[curr*2+1] {
			minHeap.arr[curr], minHeap.arr[curr*2+1] = minHeap.arr[curr*2+1], minHeap.arr[curr]
			curr = curr*2 + 1
		} else {
			break
		}
	}

	return result
}

func (maxHeap *MaxHeap) Push(val int) {
	maxHeap.arr = append(maxHeap.arr, val)
	curr := len(maxHeap.arr) - 1

	for curr >= 1 {
		if maxHeap.arr[curr] > maxHeap.arr[(curr-1)/2] {
			maxHeap.arr[curr], maxHeap.arr[(curr-1)/2] = maxHeap.arr[(curr-1)/2], maxHeap.arr[curr]
		} else {
			break
		}

		curr = (curr - 1) / 2
	}
}

func (maxHeap *MaxHeap) Pop() int {
	result := maxHeap.arr[0]
	maxHeap.arr[0], maxHeap.arr[len(maxHeap.arr)-1] = maxHeap.arr[len(maxHeap.arr)-1], maxHeap.arr[0]
	maxHeap.arr = maxHeap.arr[:len(maxHeap.arr)-1]

	curr := 0

	for curr*2+1 < len(maxHeap.arr) {
		if curr*2+2 < len(maxHeap.arr) && maxHeap.arr[curr*2+2] > maxHeap.arr[curr*2+1] && maxHeap.arr[curr*2+2] > maxHeap.arr[curr] {
			maxHeap.arr[curr], maxHeap.arr[curr*2+2] = maxHeap.arr[curr*2+2], maxHeap.arr[curr]
			curr = curr*2 + 2
		} else if maxHeap.arr[curr*2+1] > maxHeap.arr[curr] {
			maxHeap.arr[curr], maxHeap.arr[curr*2+1] = maxHeap.arr[curr*2+1], maxHeap.arr[curr]
			curr = curr*2 + 1
		} else {
			break
		}
	}

	return result
}

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.maxHeap.Push(num)
	if len(this.minHeap.arr) > 0 && this.maxHeap.arr[0] > this.minHeap.arr[0] {
		this.maxHeap.Push(this.minHeap.Pop())
	}

	for len(this.maxHeap.arr)-len(this.minHeap.arr) > 1 {
		this.minHeap.Push(this.maxHeap.Pop())
	}

	for len(this.minHeap.arr)-len(this.maxHeap.arr) > 1 {
		this.maxHeap.Push(this.minHeap.Pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap.arr) > len(this.minHeap.arr) {
		return float64(this.maxHeap.arr[0])
	} else if len(this.minHeap.arr) > len(this.maxHeap.arr) {
		return float64(this.minHeap.arr[0])
	}

	return (float64(this.maxHeap.arr[0]) + float64(this.minHeap.arr[0])) / 2
}
