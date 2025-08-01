package main

import "fmt"

func removeDuplicates(nums []int) int {
	left := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[left-1] != nums[i] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 3, 4, 4, 4, 4, 5}))
}
