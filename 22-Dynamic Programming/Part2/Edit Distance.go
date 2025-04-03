package main

import "fmt"

func main() {
	fmt.Println(minDistance("asdfg", "asefa"))
}

func minDistance(word1 string, word2 string) int {
	N, M := len(word1), len(word2)

	D := make([]int, M+1)

	for j := 0; j <= M; j++ {
		D[j] = j
	}

	for i := 1; i <= N; i++ {
		diag := D[0]
		D[0] = i
		for j := 1; j <= M; j++ {
			curr := D[j]
			D[j] = min(D[j]+1, D[j-1]+1, diag+m(word1[i-1], word2[j-1]))
			diag = curr
		}
	}
	return D[M]
}

func m(a, b byte) int {
	if a == b {
		return 0
	} else {
		return 1
	}
}
