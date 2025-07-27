package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	prev2, prev1 := 0, 1
	for i := 2; i <= n; i++ {
		prev2, prev1 = prev1, prev2+prev1
	}
	return prev1
}
