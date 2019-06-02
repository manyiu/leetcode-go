package main

import "fmt"

func toLowerCase(str string) string {
	l := ""
	for _, c := range str {
		if c >= 65 && c <= 90 {
			l = l + string(c+32)
		} else {
			l = l + string(c)
		}
	}

	return l
}

func main() {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
}
