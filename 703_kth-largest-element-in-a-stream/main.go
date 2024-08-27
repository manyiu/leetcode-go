package kthlargestelementinastream

type KthLargest struct {
	k       int
	minHeap []int
}

func Constructor(k int, nums []int) KthLargest {
	output := KthLargest{
		k:       k,
		minHeap: []int{0},
	}

	for _, v := range nums {
		output.Add(v)
	}

	return output
}

func (this *KthLargest) Add(val int) int {
	this.minHeap = append(this.minHeap, val)
	curr := len(this.minHeap) - 1

	for curr > 1 {
		if this.minHeap[curr/2] > this.minHeap[curr] {
			this.minHeap[curr/2], this.minHeap[curr] = this.minHeap[curr], this.minHeap[curr/2]
		} else {
			break
		}

		curr = curr / 2
	}

	for len(this.minHeap)-1 > this.k {
		this.Pop()
	}

	return this.minHeap[1]
}

func (this *KthLargest) Pop() {
	this.minHeap = append([]int{0, this.minHeap[len(this.minHeap)-1]}, this.minHeap[2:len(this.minHeap)-1]...)
	curr := 1

	for curr*2 < len(this.minHeap) {
		if curr*2+1 < len(this.minHeap) && this.minHeap[curr] > this.minHeap[curr*2+1] && this.minHeap[curr*2+1] < this.minHeap[curr*2] {
			this.minHeap[curr], this.minHeap[curr*2+1] = this.minHeap[curr*2+1], this.minHeap[curr]
			curr = curr*2 + 1
		} else if this.minHeap[curr] > this.minHeap[curr*2] {
			this.minHeap[curr], this.minHeap[curr*2] = this.minHeap[curr*2], this.minHeap[curr]
			curr = curr * 2
		} else {
			break
		}
	}
}
