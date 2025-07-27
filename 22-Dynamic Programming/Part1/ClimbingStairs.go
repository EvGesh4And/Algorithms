package main

func climbStairs(n int) int {
	// подъем по лестнице возможен только при n >= 1
	if n == 1 {
		return n
	}
	prev2, prev1 := 1, 2
	for i := 3; i <= n; i++ {
		prev2, prev1 = prev1, prev2+prev1
	}
	return prev1
}
