package main

func longestSubarray(nums []int) int {
	maxWindowSize := 0
	windowCountZero := 0

	begin := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {

			windowCountZero++
		}

		for windowCountZero > 1 {
			if nums[begin] == 0 {
				windowCountZero--
			}
			begin++
		}
		currWindowSize := end - begin
		if currWindowSize > maxWindowSize {
			maxWindowSize = currWindowSize
		}
	}

	return maxWindowSize
}
