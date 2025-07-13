package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	// Example usage of IPO function
	profit := findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2})
	fmt.Println(profit) // Output: 30
}

// Проект с требуемым капиталом и прибылью
type Project struct {
	capital int
	profit  int
}

// MaxHeap реализует heap.Interface по убыванию прибыли
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // max-heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]Project, n)

	for i := 0; i < n; i++ {
		projects[i] = Project{capital[i], profits[i]}
	}

	// Сортируем проекты по требуемому капиталу
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	h := &MaxHeap{}
	heap.Init(h)

	i := 0
	for k > 0 {
		// Добавляем все доступные проекты в кучу
		for i < n && projects[i].capital <= w {
			heap.Push(h, projects[i].profit)
			i++
		}

		// Нет доступных проектов
		if h.Len() == 0 {
			break
		}

		// Выполняем самый прибыльный проект
		w += heap.Pop(h).(int)
		k--
	}

	return w
}
