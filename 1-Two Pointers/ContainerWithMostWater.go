package main

import "fmt"

func maxArea(height []int) int {
	maxV := 0
	left, right := 0, len(height)-1
	for left < right {
		v := (right - left) * min(height[left], height[right])
		maxV = max(maxV, v)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxV
}

func main() {
	fmt.Println(maxArea([]int{1, 3, 0, 0, 3, 0, 0, 1}))
}
