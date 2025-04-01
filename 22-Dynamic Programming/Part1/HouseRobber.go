package main

func rob(nums []int) int {
	a, b := 0, nums[0]

	for i := 2; i <= len(nums); i++ {
		curr := max(nums[i-1]+a, b)
		a, b = b, curr
	}

	return b
}

func max(a, b int) (c int) {
	c = a
	if b > a {
		c = b
	}
	return
}
