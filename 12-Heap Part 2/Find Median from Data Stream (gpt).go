package main

import (
	"container/heap"
	"fmt"
)

func main() {
	stream := Constructor()
	stream.AddNum(1)
	stream.AddNum(2)
	fmt.Println(stream.FindMedian()) // Output: 1.5
	stream.AddNum(3)
	fmt.Println(stream.FindMedian()) // Output: 2
	stream.AddNum(10)
	stream.AddNum(6)
	fmt.Println(stream.FindMedian())
	stream.AddNum(6)
	fmt.Println(stream.FindMedian())
	stream.AddNum(6)
	fmt.Println(stream.FindMedian())
}

type MedianFinder struct {
	low  *MaxHeap
	high *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		low:  &MaxHeap{},
		high: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.low.Len() == 0 || num <= (*this.low)[0] {
		heap.Push(this.low, num)
	} else {
		heap.Push(this.high, num)
	}

	// Балансировка: разница в размере <= 1
	if this.low.Len() > this.high.Len()+1 {
		heap.Push(this.high, heap.Pop(this.low))
	} else if this.high.Len() > this.low.Len() {
		heap.Push(this.low, heap.Pop(this.high))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.low.Len() == this.high.Len() {
		return float64((*this.low)[0]+(*this.high)[0]) / 2.0
	}
	return float64((*this.low)[0])
}

// MaxHeap реализует max-heap
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinHeap реализует min-heap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
