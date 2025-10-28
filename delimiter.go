package main

import (
	"fmt"
)

func main() {
	fmt.Println("abc[", DelimiterMatch("abc[]}"))
}

func DelimiterMatch(str string) bool {
	var stack []byte
	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			stack = append(stack, str[i])
		} else if str[i] == ']' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if str[i] == '{' {
			stack = append(stack, str[i])
		} else if str[i] == '}' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println("Stack", stack)
	return len(stack) == 0
}
