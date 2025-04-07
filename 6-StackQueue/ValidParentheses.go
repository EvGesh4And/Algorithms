package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune, 0, len(s))

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
