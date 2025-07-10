package main

import "fmt"

func main() {
	// Example usage of InsertionSort
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	InsertionSort(arr)
	fmt.Println("Sorted array:", arr)
}

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
}
