func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	mask := 1
	for mask&xor == 0 {
		mask = mask << 1
	}

	a, b := 0, 0
	for _, num := range nums {
		if num&mask != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}