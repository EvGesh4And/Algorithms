package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // O(n log n)
	n := len(nums)
	var res [][]int

	for i := 0; i < n-2; i++ {
		// Для исключения работы с дубликатами
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -1 * nums[i]
		left := i + 1
		right := n - 1

		for left < right {
			curr := nums[left] + nums[right]
			if curr == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// Для исключения работы с дубликатами
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// Для исключения работы с дубликатами
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if curr > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{0, 1, 1}))
}
