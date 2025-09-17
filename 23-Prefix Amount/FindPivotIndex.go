package main

func pivotIndex(nums []int) int {

	n := len(nums)
	prefSums := make([]int, n+1)

	for i, num := range nums {
		prefSums[i+1] = prefSums[i] + num
	}

	for i := range n {
		if prefSums[n]-prefSums[i+1] == prefSums[i] {
			return i
		}
	}

	return -1
}
