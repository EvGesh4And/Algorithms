package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxS := 0

	for left < right {
		h := height[left]
		if h > height[right] {
			h = height[right]
		}
		s := (right - left) * h
		if s > maxS {
			maxS = s
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxS
}

func main() {
	fmt.Println(maxArea([]int{1, 3, 0, 0, 3, 0, 0, 1}))
}
