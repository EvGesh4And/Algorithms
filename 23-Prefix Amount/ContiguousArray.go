package main

func findMaxLength(nums []int) int {

	res := 0
	m := make(map[int]int)
	prefSum := 0

	m[0] = 0

	for i, num := range nums {
		if num == 1 {
			prefSum++
		} else {
			prefSum--
		}

		if idx, ok := m[prefSum]; ok {
			res = max(res, i+1-idx)
		} else {
			m[prefSum] = i + 1
		}
	}

	return res
}
