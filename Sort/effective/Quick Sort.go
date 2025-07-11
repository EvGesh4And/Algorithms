package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Example usage of Quick Sort
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	QuickSort(arr)
	fmt.Println("Sorted array:", arr)
}

func QuickSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	// Randomly select a pivot to avoid worst-case performance on sorted input
	rnd := rand.Intn(n)
	if rnd != n-1 {
		arr[rnd], arr[n-1] = arr[n-1], arr[rnd]
	}

	pivot := arr[n-1]
	idx := 0
	for i := 0; i < n-1; i++ {
		if arr[i] < pivot {
			arr[idx], arr[i] = arr[i], arr[idx]
			idx++
		}
	}
	arr[idx], arr[n-1] = arr[n-1], arr[idx]
	QuickSort(arr[:idx])
	QuickSort(arr[idx+1:])
}
