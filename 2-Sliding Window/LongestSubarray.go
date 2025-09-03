func longestSubarray(nums []int) int {
	maxLen := 0
	zeros := 0
	left := 0
	for i := range nums {
		if nums[i] == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		maxLen = max(maxLen, i-left)
	}

	return maxLen
}