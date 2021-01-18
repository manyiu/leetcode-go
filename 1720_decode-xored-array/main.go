package main

func decode(encoded []int, first int) []int {
    r := make([]int, len(encoded) + 1)
	r[0] = first

	for i, v := range encoded {
		r[i + 1] = r[i] ^ v
	}

	return r
}

func main() {}