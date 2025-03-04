package main

func climbStairs(n int) int {
	a, b := 1, 1

	for i := 2; i <= n; i++ {
		curr := a + b
		a, b = b, curr
	}

	return b
}
