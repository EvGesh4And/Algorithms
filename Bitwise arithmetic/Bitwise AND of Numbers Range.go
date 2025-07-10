package main

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left != right {
		shift++
		left >>= 1
		right >>= 1
	}
	return left << shift
}
