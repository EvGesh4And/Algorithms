func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// Инициализируем dp для первой строки
	dp := make([]int, n)
	dp[0] = grid[0][0]

	// Обработка первой строки
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	// Обработка остальных строк
	for i := 1; i < m; i++ {
		// Обновляем dp[0] для первого столбца
		dp[0] = dp[0] + grid[i][0]
		// Обновляем остальные элементы dp
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}

	return dp[n-1]
}
