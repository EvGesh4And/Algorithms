package main

import "strings"

func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]byte, n)
	for i := range board {
		board[i] = []byte(strings.Repeat(".", n))
	}

	columns := make(map[int]bool)
	diag1 := make(map[int]bool) // row - col
	diag2 := make(map[int]bool) // row + col

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var solution []string
			for i := 0; i < n; i++ {
				solution = append(solution, string(board[i]))
			}
			result = append(result, solution)
			return
		}

		for col := 0; col < n; col++ {
			if columns[col] || diag1[row-col] || diag2[row+col] {
				continue
			}

			board[row][col] = 'Q'
			columns[col] = true
			diag1[row-col] = true
			diag2[row+col] = true

			backtrack(row + 1)

			board[row][col] = '.'
			delete(columns, col)
			delete(diag1, row-col)
			delete(diag2, row+col)
		}
	}

	backtrack(0)
	return result
}
