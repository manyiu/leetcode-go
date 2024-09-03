package designbrowserhistory

type node struct {
	url  string
	prev *node
	next *node
}

type BrowserHistory struct {
	current *node
	head    *node
	tail    *node
}

func Constructor(homepage string) BrowserHistory {
	head := &node{}
	tail := &node{}
	current := &node{
		url:  homepage,
		prev: head,
		next: tail,
	}
	head.next, tail.prev = current, current

	return BrowserHistory{
		current: current,
		head:    head,
		tail:    tail,
	}
}

func (this *BrowserHistory) Visit(url string) {
	newNode := &node{
		url:  url,
		prev: this.current,
		next: this.tail,
	}

	this.current.next, this.tail.prev, this.current = newNode, newNode, newNode
}

func (this *BrowserHistory) Back(steps int) string {
	curr := this.current

	for i := steps; i > 0 && curr.prev != this.head; i-- {
		curr = curr.prev
	}

	this.current = curr

	return curr.url
}

func (this *BrowserHistory) Forward(steps int) string {
	curr := this.current

	for i := steps; i > 0 && curr.next != this.tail; i-- {
		curr = curr.next
	}

	this.current = curr

	return curr.url
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
