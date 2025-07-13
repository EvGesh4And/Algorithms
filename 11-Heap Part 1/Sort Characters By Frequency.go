package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Abbcccadddd"
	res := frequencySort(s)
	fmt.Println("res: ", res)
}

func frequencySort(s string) string {
	m := make(map[rune]int)

	for _, rn := range s {
		m[rn]++
	}

	n := len(m)

	heap := make([]pair, 0, n)

	for r, cnt := range m {
		heap = append(heap, pair{r, cnt})
	}

	fmt.Println(heap)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(heap, n, i)
	}

	for i := n - 1; i > 0; i-- {
		heap[0], heap[i] = heap[i], heap[0]
		heapify(heap, i, 0)
	}
	var builder strings.Builder

	for i := n - 1; i >= 0; i-- {
		for _ = range heap[i].cnt {
			builder.WriteRune(heap[i].r)
		}
	}
	return builder.String()
}

type pair struct {
	r   rune
	cnt int
}

func heapify(arr []pair, n, i int) {
	idxMax := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[idxMax].cnt < arr[left].cnt {
		idxMax = left
	}
	if right < n && arr[idxMax].cnt < arr[right].cnt {
		idxMax = right
	}
	if i != idxMax {
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
		heapify(arr, n, idxMax)
	}
}
