package main

func minSubArrayLen(target int, nums []int) int {
	begin := 0
	windowSum := 0
	minWindowSize := len(nums) + 1

	for end := 0; end < len(nums); end++ {
		windowSum += nums[end]

		// Когда сумма окна больше или равна target
		for windowSum >= target {
			currWindowSize := end - begin + 1
			if currWindowSize < minWindowSize {
				minWindowSize = currWindowSize
			}
			windowSum -= nums[begin]
			begin++
		}
	}

	if minWindowSize == len(nums)+1 {
		return 0 // Если не найдено подходящее окно
	}
	return minWindowSize
}
