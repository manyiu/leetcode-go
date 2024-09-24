package designaddandsearchwordsdatastructure

type WordDictionary struct {
	children [26]*WordDictionary
	end      bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this

	for _, char := range word {
		if curr.children[char-'a'] == nil {
			curr.children[char-'a'] = &WordDictionary{}
		}

		curr = curr.children[char-'a']
	}

	curr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	for index, char := range word {
		if char == '.' {
			for _, child := range this.children {
				if child != nil && child.Search(word[index+1:]) {
					return true
				}
			}

			return false
		} else {
			if this.children[char-'a'] != nil {
				return this.children[char-'a'].Search(word[index+1:])
			} else {
				return false
			}
		}
	}

	return this.end
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
