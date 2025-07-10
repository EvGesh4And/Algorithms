package main

import "fmt"

func main() {
	// Example usage of BubbleSort
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	BubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap if the element found is greater
				swapped = true
			}
		}
		if !swapped {
			break // If no two elements were swapped, the array is sorted
		}
	}
}
