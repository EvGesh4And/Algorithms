package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		curr := numbers[left] + numbers[right]
		if curr == target {
			return []int{left + 1, right + 1}
		}
		if curr > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}
