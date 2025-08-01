package main

func totalFruit(fruits []int) int {
	fruitMap := make(map[int]int)
	maxWindowSize := 0

	begin := 0

	for end := 0; end < len(fruits); end++ {
		fruitMap[fruits[end]]++

		for len(fruitMap) > 2 {
			fruitMap[fruits[begin]]--
			if fruitMap[fruits[begin]] == 0 {
				delete(fruitMap, fruits[begin])
			}
			begin++
		}
		maxWindowSize = max(maxWindowSize, end-begin+1)
	}
	return maxWindowSize
}
