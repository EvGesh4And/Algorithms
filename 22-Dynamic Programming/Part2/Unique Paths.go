package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	dp := make([]int, n)

	for j := range n {
		dp[j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
