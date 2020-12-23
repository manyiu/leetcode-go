package main

func maximumWealth(accounts [][]int) int {
	r := 0

    for _, c := range accounts {
		w := 0

		for _, b := range c {
			w += b
		}

		if w > r {
			r = w
		}
	}

	return r
}

func main() {}