package main

func checkSubarraySum(nums []int, k int) bool {
	remFirstIdx := map[int]int{0: -1}
	sum := 0

	for i, v := range nums {
		sum += v
		rem := sum % k

		if j, ok := remFirstIdx[rem]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			remFirstIdx[rem] = i
		}
	}
	return false
}
