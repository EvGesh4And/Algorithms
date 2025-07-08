package main

func hammingWeight(n int) int {
	var cnt int
	for n > 0 {
		cnt++
		n &= (n - 1)
	}
	return cnt
}
