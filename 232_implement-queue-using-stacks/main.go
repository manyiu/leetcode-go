package implementqueueusingstacks

type Stack []int

type MyQueue struct {
	firstStack  Stack
	secondStack Stack
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return x
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.firstStack.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.secondStack.Empty() {
		for !q.firstStack.Empty() {
			q.secondStack.Push(q.firstStack.Pop())
		}
	}

	return q.secondStack.Pop()
}

func (q *MyQueue) Peek() int {
	if q.secondStack.Empty() {
		for !q.firstStack.Empty() {
			q.secondStack.Push(q.firstStack.Pop())
		}
	}

	return q.secondStack.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.firstStack.Empty() && q.secondStack.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
