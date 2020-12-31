package main

func minPartitions(n string) int {
    r := '0'

	for _, v := range n {
		if v > r {
			r = v
		}
	}

	return int(r - '0')
}

func main() {}
