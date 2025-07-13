package main

import (
	"fmt"
)

func main() {
	// Example usage of IPO function
	profit := findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2})
	fmt.Println(profit) // Output: 30
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var heap []int

	for i, cap := range capital {
		if cap <= w {
			heap = append(heap, profits[i])
		}
	}
	n := len(heap)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(heap, n, i)
	}

	cnt := 0
	for cnt < k {
		prevW := w
		if len(heap) == 0 {
			break
		}
		w += heap[0]
		if n > 1 {
			heap[0] = heap[n-1]
			heap = heap[:n-1]
			n--
			heapify(heap, n, 0)
		} else {
			heap = []int{}
			n = 0
		}
		cnt++
		for i, cap := range capital {
			if cap <= w && cap > prevW {
				heap = append(heap, profits[i])
				shiftUp(heap, n)
				n++
			}
		}
	}
	return w
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if i != largest {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func shiftUp(arr []int, i int) {
	parent := (i - 1) / 2
	if parent >= 0 && arr[parent] < arr[i] {
		arr[parent], arr[i] = arr[i], arr[parent]
		shiftUp(arr, parent)
	}
}
