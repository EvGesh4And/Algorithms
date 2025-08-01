package main

func moveZeroes(nums []int) {
	left := 0
	// Проходим по всем элементам
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// Меняем местами nums[i] и nums[left]
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
}
