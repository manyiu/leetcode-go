package designlinkedlist

type Node struct {
	val  int
	prev *Node
	next *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
}

func Constructor() MyLinkedList {
	head := &Node{
		val: 1001,
	}

	tail := &Node{
		val: 1001,
	}

	head.next = tail
	tail.prev = head

	return MyLinkedList{
		head: head,
		tail: tail,
	}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.head

	for i := 0; cur != nil && i <= index; i++ {
		cur = cur.next
	}

	if cur != nil && cur != this.tail {
		return cur.val
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		val:  val,
		prev: this.head,
		next: this.head.next,
	}

	this.head.next, this.head.next.prev = newNode, newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		val:  val,
		prev: this.tail.prev,
		next: this.tail,
	}

	this.tail.prev, this.tail.prev.next = newNode, newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.head

	for i := 0; cur != nil && i <= index; i++ {
		cur = cur.next
	}

	if cur == nil {
		return
	}

	newNode := &Node{
		val:  val,
		prev: cur.prev,
		next: cur,
	}

	cur.prev, cur.prev.next = newNode, newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.head

	for i := 0; cur != nil && i <= index; i++ {
		cur = cur.next
	}

	if cur == this.tail || cur == nil {
		return
	}

	cur.prev.next, cur.next.prev = cur.next, cur.prev
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
