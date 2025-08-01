package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	i := 0
	n := len(intervals)

	// 1. Добавляем все, которые до newInterval
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// 2. Объединяем все, что пересекаются
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)

	// 3. Добавляем оставшиеся
	for i < n {
		res = append(res, intervals[i])
		i++
	}

	return res
}
