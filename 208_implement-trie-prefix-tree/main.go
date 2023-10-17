package implementtrieprefixtree

type Trie struct {
	children [26]*Trie
	end      bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	c := this

	for _, ch := range word {
		a := ch - 'a'

		if c.children[a] == nil {
			c.children[a] = &Trie{}
		}

		c = c.children[a]
	}

	c.end = true
}

func (this *Trie) Search(word string) bool {
	c := this

	for _, ch := range word {
		a := ch - 'a'

		if c.children[a] != nil {
			c = c.children[a]
		} else {
			return false
		}
	}

	return c.end
}

func (this *Trie) StartsWith(prefix string) bool {
	c := this

	for _, ch := range prefix {
		a := ch - 'a'

		if c.children[a] != nil {
			c = c.children[a]
		} else {
			return false
		}
	}

	return true
}
