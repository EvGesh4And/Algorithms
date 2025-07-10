package main

import "fmt"

func main() {
	// Example usage of InsertionSort
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	SelectionSort(arr)
	fmt.Println("Sorted array:", arr)
}

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if arr[idx] > arr[j] {
				idx = j
			}
		}
		if i != idx {
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
}
