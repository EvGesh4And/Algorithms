package main

import "fmt"

func maxArea(height []int) int {
	maxV := 0
	left, right := 0, len(height)-1
	for left < right {
		width := right - left
		v := min(height[left], height[right]) * width
		if v > maxV {
			maxV = v
		}
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
