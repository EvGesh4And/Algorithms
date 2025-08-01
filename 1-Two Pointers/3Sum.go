package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		target := -nums[i]
		left, right := i+1, n-1

		for left < right {
			sum := nums[left] + nums[right]
			switch {
			case sum < target:
				left++
			case sum > target:
				right--
			default:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{0, 1, 1}))
}
