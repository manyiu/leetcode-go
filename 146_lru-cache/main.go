package lrucache

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	length   int
	hash     map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		length:   0,
		hash:     map[int]*Node{},
		head:     head,
		tail:     tail,
	}
}

func (node *Node) Unlink() {
	node.prev.next, node.next.prev = node.next, node.prev
}

func (this *LRUCache) Enqueue(node *Node) {
	node.prev, node.next = this.tail.prev, this.tail
	this.tail.prev.next, this.tail.prev = node, node
}

func (this *LRUCache) Dequeue() {
	delete(this.hash, this.head.next.key)
	this.head.next.Unlink()
}

func (this *LRUCache) Get(key int) int {
	existedNode, existed := this.hash[key]

	if !existed {
		return -1
	}

	existedNode.Unlink()
	this.Enqueue(existedNode)

	return existedNode.val
}

func (this *LRUCache) Put(key int, value int) {
	node, existed := this.hash[key]

	if !existed {
		this.length++
		node = &Node{
			key: key,
			val: value,
		}
	} else {
		node.val = value
		node.Unlink()
	}

	this.Enqueue(node)
	this.hash[key] = node

	if this.length > this.capacity {
		this.Dequeue()
		this.length--
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
