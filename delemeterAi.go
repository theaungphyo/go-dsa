package main

import "fmt"

func main() {
	fmt.Println("abc[]{}", DelimeterMatch("abc[]{}"))  // ✅ true
	fmt.Println("abc[}", DelimeterMatch("abc[}"))      // ❌ false
	fmt.Println("abc{[}]}", DelimeterMatch("abc{[}]}")) // ❌ false
	fmt.Println("abc", DelimeterMatch("abc"))           // ✅ true
}

func DelimeterMatch(str string) bool {
	pairs := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}

	var stack []byte

	for i := 0; i < len(str); i++ {
		ch := str[i]

		switch ch {
		case '[', '{', '(':
			stack = append(stack, ch)

		case ']', '}', ')':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
