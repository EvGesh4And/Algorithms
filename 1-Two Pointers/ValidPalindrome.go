package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !(unicode.IsDigit(rune(s[left])) || unicode.IsLetter(rune(s[left]))) {
			left++
			continue
		}
		for !(unicode.IsDigit(rune(s[right])) || unicode.IsLetter(rune(s[right]))) {
			right--
			continue
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "Abdd-*1fq&*^qf1dba"
	fmt.Println(isPalindrome(s))
}
