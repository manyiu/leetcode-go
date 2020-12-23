package main

func interpret(command string) string {
	r := []byte{}
	
	for i := 0; i < len(command); i++ {
		if command[i] == '(' {
			if i + 1 < len(command) && command[i + 1] == ')' {
				r = append(r, 'o')
				i++
			} else if i + 3 < len(command) && command[i + 1] == 'a' && command[i+2] == 'l' && command[i+3] == ')' {
				r = append(r, 'a', 'l')
				i += 3
			}
		} else {
			r = append(r, byte(command[i]))
		}
	}

	return string(r)
}

func main() {}