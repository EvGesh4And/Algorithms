package main

import (
	"fmt"
	"math"
)

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
