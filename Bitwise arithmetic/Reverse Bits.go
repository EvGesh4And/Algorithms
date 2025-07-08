package main

import (
	"fmt"
	"math"
)

func main() {
	// Example usage
	num := uint32(234)         // Input number
	fmt.Printf("%032b\n", num) // 32-bit binary representation
	result := reverseBits(num)
	fmt.Printf("%032b\n", result) // Output: 964176192
}

func reverseBitsBest(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		bit := num & 1 // взять младший бит
		result <<= 1   // сдвинуть результат влево
		result |= bit  // вставить бит в младший разряд
		num >>= 1      // сдвинуть num вправо
	}
	return result
}

func reverseBits(num uint32) uint32 {
	left, right := 31, 0
	var leftBit uint32
	var rightBit uint32
	var maxUint32 uint32 = math.MaxUint32

	for right < 16 {
		rightBit = (num &^ (maxUint32 &^ (1 << left))) >> (left - right)
		leftBit = (num &^ (maxUint32 &^ (1 << right))) << (left - right)
		num = (num &^ (1 << left)) | leftBit
		num = (num &^ (1 << right)) | rightBit
		left--
		right++
	}

	return num
}
