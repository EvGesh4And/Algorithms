package main

func rob(nums []int) int {
	prev2, prev1 := 0, 0
	for i := range len(nums) {
		prev2, prev1 = prev1, max(prev2+nums[i], prev1)
	}
	return prev1
}
