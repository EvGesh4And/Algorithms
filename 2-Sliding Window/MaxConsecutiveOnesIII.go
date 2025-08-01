package main

func longestOnes(nums []int, k int) int {
	windowCountZero := 0
	maxWindowSize := 0
	begin := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			windowCountZero++
		}

		for windowCountZero > k {
			if nums[begin] == 0 {
				windowCountZero--
			}
			begin++
		}

		maxWindowSize = max(maxWindowSize, end-begin+1)
	}

	return maxWindowSize
}
