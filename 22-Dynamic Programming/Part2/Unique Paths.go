package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	matrix[0][0] = 1

	for i, line := range matrix {
		for j, _ := range line {
			if j != 0 {
				line[j] += line[j-1]
			}
			if i != 0 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}

	return matrix[m-1][n-1]
}
