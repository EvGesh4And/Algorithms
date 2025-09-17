package main

func runningSum(nums []int) []int {
	var prefSums []int
	prefSum := 0
	for _, num := range nums {
		prefSum += num
		prefSums = append(prefSums, prefSum)
	}

	return prefSums
}
