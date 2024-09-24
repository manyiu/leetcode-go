package prefixandsuffixsearch

type Trie struct {
	children [26]*Trie
	indexes  []int
}

type WordFilter struct {
	prefix *Trie
	suffix *Trie
}

func Constructor(words []string) WordFilter {
	prefix := &Trie{}
	suffix := &Trie{}

	for index, word := range words {
		curr := prefix

		for i := 0; i < len(word); i++ {
			if curr.children[word[i]-'a'] == nil {
				curr.children[word[i]-'a'] = &Trie{
					indexes: []int{index},
				}
			} else {
				curr.children[word[i]-'a'].indexes = append(curr.children[word[i]-'a'].indexes, index)
			}
			curr = curr.children[word[i]-'a']
		}

		curr = suffix

		for j := len(word) - 1; j >= 0; j-- {
			if curr.children[word[j]-'a'] == nil {
				curr.children[word[j]-'a'] = &Trie{
					indexes: []int{index},
				}
			} else {
				curr.children[word[j]-'a'].indexes = append(curr.children[word[j]-'a'].indexes, index)
			}
			curr = curr.children[word[j]-'a']
		}

	}

	return WordFilter{
		prefix,
		suffix,
	}
}

func (this *WordFilter) F(pref string, suff string) int {
	curr := this.prefix

	for _, p := range pref {
		if curr.children[p-'a'] != nil {
			curr = curr.children[p-'a']
		} else {
			return -1
		}
	}

	prefIndexes := curr.indexes

	curr = this.suffix

	for k := len(suff) - 1; k >= 0; k-- {
		if curr.children[suff[k]-'a'] != nil {
			curr = curr.children[suff[k]-'a']
		} else {
			return -1
		}
	}

	suffIndexes := curr.indexes

	i := len(prefIndexes) - 1
	j := len(suffIndexes) - 1

	for i >= 0 && j >= 0 {
		if prefIndexes[i] == suffIndexes[j] {
			return prefIndexes[i]
		}

		if prefIndexes[i] > suffIndexes[j] {
			i--
		} else {
			j--
		}
	}

	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
