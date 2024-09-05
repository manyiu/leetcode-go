package implementstackusingqueues

type Queue []int

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}

func (q *Queue) Top() int {
	return (*q)[0]
}

func (q *Queue) Pop() int {
	v := (*q)[0]
	*q = (*q)[1:]

	return v
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

type MyStack struct {
	queue Queue
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue.Push(x)
}

func (this *MyStack) Pop() int {
	temp := Queue{}

	for len(this.queue) > 1 {
		temp.Push(this.queue.Pop())
	}

	popValue := this.queue.Pop()

	this.queue = temp

	return popValue
}

func (this *MyStack) Top() int {
	temp := Queue{}

	for len(this.queue) > 1 {
		temp.Push(this.queue.Pop())
	}

	topValue := this.queue.Top()

	temp.Push(this.queue.Pop())

	this.queue = temp

	return topValue
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
