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
			minWindowSize = min(minWindowSize, currWindowSize)
			windowSum -= nums[begin]
			begin++
		}
	}

	if minWindowSize == len(nums)+1 {
		return 0 // Если не найдено подходящее окно
	}
	return minWindowSize
}

// Повторил решение не глядя в тек. реализацию -- ВАУ!
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minL := n + 1
	sum := 0
	pL := 0
	for i := range len(nums) {
		sum += nums[i]
		for sum >= target {
			minL = min(minL, i-pL+1)
			sum -= nums[pL]
			pL++
		}
	}
	if minL != n+1 {
		return minL
	}
	return 0
}
