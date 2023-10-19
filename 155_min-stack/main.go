package minstack

type MinStack struct {
	elements []int
	minimums []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.minimums) == 0 || val <= this.GetMin() {
		this.minimums = append(this.minimums, val)
	}

	this.elements = append(this.elements, val)
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.minimums = this.minimums[:len(this.minimums)-1]
	}

	this.elements = this.elements[:len(this.elements)-1]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.minimums[len(this.minimums)-1]
}
