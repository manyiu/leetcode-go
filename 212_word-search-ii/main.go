package wordsearchii

type trie struct {
	children [26]*trie
	isWord   bool
}

func (t *trie) add(word string) {
	curr := t

	for _, char := range word {
		if curr.children[char-'a'] == nil {
			curr.children[char-'a'] = &trie{}
		}

		curr = curr.children[char-'a']
	}

	curr.isWord = true
}

func dfs(board [][]byte, i, j int, visited [][]bool, node *trie, word string, result *[]string) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || visited[i][j] || node.children[board[i][j]-'a'] == nil {
		return
	}

	newWord := word + string(board[i][j])
	newNode := node.children[board[i][j]-'a']

	if newNode.isWord {
		*result = append(*result, newWord)
		newNode.isWord = false
	}

	visited[i][j] = true

	dfs(board, i, j-1, visited, newNode, newWord, result)
	dfs(board, i, j+1, visited, newNode, newWord, result)
	dfs(board, i-1, j, visited, newNode, newWord, result)
	dfs(board, i+1, j, visited, newNode, newWord, result)

	visited[i][j] = false
}

func findWords(board [][]byte, words []string) []string {
	wordTrie := &trie{}

	for _, word := range words {
		wordTrie.add(word)
	}

	visited := make([][]bool, len(board))

	for k := 0; k < len(board); k++ {
		row := make([]bool, len(board[k]))
		visited[k] = row
	}

	result := []string{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, i, j, visited, wordTrie, "", &result)
		}
	}

	return result
}
