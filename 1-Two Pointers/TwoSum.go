package main

import "fmt"

func twoSum(nums []int, target int) []int {
	remMap := make(map[int]int)
	for i, v := range nums {
		rem := target - v
		if idx, ok := remMap[rem]; ok {
			return []int{idx, i}
		} else {
			remMap[v] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
