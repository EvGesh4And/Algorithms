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

// Как вариант
func isValid(s string) bool {
	left := map[rune]struct{}{'(': struct{}{}, '[': struct{}{}, '{': struct{}{}}
	right := map[rune]rune{']': '[', ')': '(', '}': '{'}
	var stack []rune
	for _, r := range s {
		if _, ok := left[r]; ok {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != right[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
