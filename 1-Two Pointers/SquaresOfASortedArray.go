package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for i := len(nums) - 1; i >= 0; i-- {
		sqLeft := nums[left] * nums[left]
		sqRight := nums[right] * nums[right]
		if sqLeft >= sqRight {
			res[i] = sqLeft
			left++
		} else {
			res[i] = sqRight
			right--
		}
	}
	return res
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
