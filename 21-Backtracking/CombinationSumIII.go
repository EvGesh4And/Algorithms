package main

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var path []int

	var backtrack func(start int, sum int)
	backtrack = func(start int, sum int) {
		if len(path) == k {
			if sum == n {
				comb := make([]int, k)
				copy(comb, path)
				result = append(result, comb)
			}
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n {
				break // отсечение лишнего
			}
			path = append(path, i)
			backtrack(i+1, sum+i)
			path = path[:len(path)-1] // откат
		}
	}

	backtrack(1, 0)
	return result
}
