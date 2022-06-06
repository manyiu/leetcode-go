package validparentheses

func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, r := range s {
		if _, ok := m[r]; ok {
			stack = append(stack, r)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if r != m[stack[len(stack)-1]] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
