package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	pLeft, pRight := 0, len(s)-1

	for pLeft < pRight {
		if !(unicode.IsDigit(rune(s[pLeft])) || unicode.IsLetter(rune(s[pLeft]))) {
			pLeft++
			continue
		}
		if !(unicode.IsDigit(rune(s[pRight])) || unicode.IsLetter(rune(s[pRight]))) {
			pRight--
			continue
		}
		if unicode.ToLower(rune(s[pLeft])) != unicode.ToLower(rune(s[pRight])) {
			return false
		}
		pLeft++
		pRight--
	}
	return true
}

func main() {
	s := "Abdd-*1fq&*^qf1dba"
	fmt.Println(isPalindrome(s))
}
