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
	low   *MaxHeap
	hight *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		low:   &MaxHeap{},
		hight: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.low.Len() == 0 {
		this.addLow(num)
		return
	}
	if this.hight.Len() == 0 {
		this.addLow(num)
		return
	}
	if (*this.low)[0] > num {
		this.addLow(num)
		return
	}
	this.addHight(num)
}

func (this *MedianFinder) addLow(num int) {
	heap.Push(this.low, num)
	if this.low.Len()-this.hight.Len() > 1 {
		maxLow := heap.Pop(this.low).(int)
		heap.Push(this.hight, maxLow)
	}
}

func (this *MedianFinder) addHight(num int) {
	heap.Push(this.hight, num)
	if this.hight.Len()-this.low.Len() > 1 {
		minHight := heap.Pop(this.hight).(int)
		heap.Push(this.low, minHight)
	}
}

func (this *MedianFinder) FindMedian() float64 {

	if this.hight.Len() == this.low.Len() {
		return float64((*this.hight)[0]+(*this.low)[0]) / 2
	}
	if this.hight.Len() > this.low.Len() {
		return float64((*this.hight)[0])
	}
	return float64((*this.low)[0])
}

type MaxHeap []int

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m *MaxHeap) Push(x any)        { *m = append(*m, x.(int)) }
func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

type MinHeap []int

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MinHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m *MinHeap) Push(x any)        { *m = append(*m, x.(int)) }
func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
