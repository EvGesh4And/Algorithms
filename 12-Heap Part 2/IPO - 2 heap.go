package main

import (
	"container/heap"
	"fmt"
)

func main() {
	profit := findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2})
	fmt.Println(profit) // Output: 6
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	hMin := &minCapital{}
	hMax := &maxProfit{}

	for i := range len(profits) {
		heap.Push(hMin, project{profits[i], capital[i]})
	}

	for i := 0; i < k; i++ {
		for hMin.Len() > 0 && (*hMin)[0].capital <= w {
			heap.Push(hMax, heap.Pop(hMin))
		}
		if hMax.Len() > 0 {
			w += heap.Pop(hMax).(project).profit
		} else {
			break
		}
	}

	return w
}

type project struct {
	profit  int
	capital int
}

type minCapital []project

func (m minCapital) Len() int           { return len(m) }
func (m minCapital) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m minCapital) Less(i, j int) bool { return m[i].capital < m[j].capital }
func (m *minCapital) Push(x any)        { *m = append(*m, x.(project)) }
func (h *minCapital) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type maxProfit []project

func (m maxProfit) Len() int           { return len(m) }
func (m maxProfit) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m maxProfit) Less(i, j int) bool { return m[i].profit > m[j].profit }
func (m *maxProfit) Push(x any)        { *m = append(*m, x.(project)) }
func (h *maxProfit) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
