package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(addBinary("101", "101110"))
}

func addBinary(a string, b string) string {
	var result bytes.Buffer
	carry := 0
	i, j := len(a)-1, len(b)-1
ы
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result.WriteByte(byte(sum%2 + '0'))
		carry = sum / 2
	}

	// Reverse the result since we built it backwards
	runes := []rune(result.String())
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	return string(runes)

}
