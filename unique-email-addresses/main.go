package main

import "fmt"

func numUniqueEmails(emails []string) int {
	m := map[string]bool{}

	for _, email := range emails {
		t := ""
		s := false
		d := false

		for _, c := range email {
			if c == '.' && d == false {
				continue
			} else if c == '+' && d == false {
				s = true
			} else if c == '@' {
				s = false
				d = true
			}

			if s != true {
				t += string(c)
			}
		}

		m[t] = true
	}

	return len(m)
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
