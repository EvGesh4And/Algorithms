package main

func subarraySum(nums []int, k int) int {
	count := 0
	m := make(map[int]int) // префиксная сумма -> сколько раз встречалась
	prefSum := 0

	m[0] = 1 // пустой префикс

	for _, num := range nums {
		prefSum += num
		// сколько было префиксов с суммой (prefSum - k)
		count += m[prefSum-k]
		m[prefSum]++
	}
	return count
}
