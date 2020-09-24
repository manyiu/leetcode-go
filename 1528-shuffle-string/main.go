package main

func restoreString(s string, indices []int) string {
	r := make([]rune, len(s))
	
	for i, v := range s {
		r[indices[i]] = v
	}

	return string(r)
}

func main() {}