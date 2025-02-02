package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	curr := 0 // Индекс, куда записываем следующий уникальный элемент

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[curr] {
			curr++
			nums[curr] = nums[i]
		}
	}
	return curr + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 3, 4, 4, 4, 4, 5}))
}
