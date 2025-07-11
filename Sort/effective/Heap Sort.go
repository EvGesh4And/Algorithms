package main

import "fmt"

func main() {
	// Example usage
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	heapSort(arr)
	fmt.Println("Sorted array:", arr)
}

func heapSort(arr []int) {
	n := len(arr)

	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// One by one extract elements from heap
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i] // swap
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, i int) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // left = 2*i + 1
	right := 2*i + 2 // right = 2*i + 2

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // swap
		heapify(arr, n, largest)
	}
}
