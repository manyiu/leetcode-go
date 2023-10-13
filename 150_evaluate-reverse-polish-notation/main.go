package evaluatereversepolishnotation

import "strconv"

func isOperator(s string) bool {
	switch s {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	}

	return false
}

func evalRPN(tokens []string) int {
	s := []int{}

	for _, token := range tokens {
		if isOperator(token) {
			a := s[len(s)-2]
			b := s[len(s)-1]
			s = s[:len(s)-2]

			switch token {
			case "+":
				s = append(s, a+b)
			case "-":
				s = append(s, a-b)
			case "*":
				s = append(s, a*b)
			case "/":
				s = append(s, a/b)
			}
		} else {
			i, _ := strconv.Atoi(token)
			s = append(s, i)
		}
	}

	return s[0]
}
