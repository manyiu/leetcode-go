package main

func maxDepth(s string) int {
	t := 0
	d := 0

	for _, v := range s {
		if v == '(' {
			t++

			if t > d {
				d = t
			}
		} else if v == ')' {
			t--
		}
	}

	return d
}

func main() {}
