package main

func subsets(nums []int) [][]int {
	var result [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// добавляем копию текущего подмножества
		subset := make([]int, len(path))
		copy(subset, path)
		result = append(result, subset)

		// итерируем по оставшимся элементам
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)          // рекурсивно добавляем следующие элементы
			path = path[:len(path)-1] // откат
		}
	}

	backtrack(0)
	return result
}
