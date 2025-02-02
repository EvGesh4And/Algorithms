package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) == unicode.ToLower(rune(s[right])) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "Abdd-*1fq&*^qf1dba"
	fmt.Println(isPalindrome(s))
}
