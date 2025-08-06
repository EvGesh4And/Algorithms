package main

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	var used = make([]bool, len(nums))

	var backtrack func()

	backtrack = func() {
		if len(path) == len(nums) {
			perm := make([]int, len(path))
			copy(perm, path)
			result = append(result, perm)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}
