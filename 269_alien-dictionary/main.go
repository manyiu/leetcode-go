package aliendictionary

func dfs(adj [][]byte, curr int, visit, path []bool, output *[]byte) bool {
	if path[curr] {
		return false
	}

	if visit[curr] {
		return true
	}

	path[curr] = true
	visit[curr] = true

	for i := 0; i < len(adj[curr]); i++ {
		if !dfs(adj, int(adj[curr][i]), visit, path, output) {
			return false
		}
	}

	*output = append(*output, byte(curr)+'a')
	path[curr] = false

	return true
}

func foreignDictionary(words []string) string {
	adj := make([][]byte, 26)

	for i := 0; i < len(words)-1; i++ {
		wordA, wordB := words[i], words[i+1]
		minLen := min(len(wordA), len(wordB))

		if len(wordA) > len(wordB) && wordA[:minLen] == wordB[:minLen] {
			return ""
		}

	CharLoop:
		for j := 0; j < minLen; j++ {
			if wordA[j] != wordB[j] {
				adj[wordA[j]-'a'] = append(adj[wordA[j]-'a'], wordB[j]-'a')
				break CharLoop
			}
		}
	}

	visit := make([]bool, 26)
	path := make([]bool, 26)
	output := []byte{}

	for i := 0; i < 26; i++ {
		if len(adj[i]) > 0 {
			if !dfs(adj, i, visit, path, &output) {
				return ""
			}
		}
	}

	return Reverse(string(output))
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
