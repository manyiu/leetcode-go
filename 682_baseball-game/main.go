package baseballgame

import "strconv"

func calPoints(operations []string) int {
	t := []int{}

	for _, v := range operations {
		switch v {
		case "+":
			t = append(t, t[len(t)-1]+t[len(t)-2])
		case "D":
			t = append(t, t[len(t)-1]*2)
		case "C":
			t = t[:len(t)-1]
		default:
			i, _ := strconv.Atoi(v)
			t = append(t, i)
		}
	}

	s := 0

	for _, v := range t {
		s += v
	}

	return s
}
