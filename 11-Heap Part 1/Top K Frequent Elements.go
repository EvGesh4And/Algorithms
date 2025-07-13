package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example usage
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Println(result) // Output: [2, 1]
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	h := &IntHeap{}
	heap.Init(h)

	for num, cnt := range m {
		heap.Push(h, []int{num, cnt})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res []int

	for _, val := range *h {
		res = append(res, val[0])
	}

	return res
}

type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
