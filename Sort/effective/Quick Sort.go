package main

import "fmt"

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
