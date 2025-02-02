package main

// строка s = [a, b, c, d, e]
//             ^           ^
// 0. берем два указателя слева и справа
// (left = 0 и right = len(s)-1)
// 1. проверяем, что left < right
// 2. меняем буквы
// 3. сдвигаем указатели к центру left++, right--
// 4. идем к шагу 1

// И ву-а-ля, мы перевернули строку!
// Тут был использован вариант *A*, подхода **Two Pointers**

import (
	"fmt"
)

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	s := []byte{'a', 'b', 'c', 'd', 'e'}
	reverseString(s)
	fmt.Println(string(s))
}
