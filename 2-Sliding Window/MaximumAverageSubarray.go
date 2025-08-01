package main

import (
	"fmt"
	"math"
)

// Новое решение (более лакончиное)
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}

func findMaxAverage(nums []int, k int) float64 {
	if k == 0 {
		return math.NaN()
	}

	begin := 0
	windowSum := 0
	maxSum := math.MinInt32

	for end := 0; end < len(nums); end++ {
		windowSum += nums[end]

		if end-begin+1 == k {
			// Обновляем максимальную сумму
			if maxSum < windowSum {
				maxSum = windowSum
			}
			// Двигаем окно
			windowSum -= nums[begin]
			begin++
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
