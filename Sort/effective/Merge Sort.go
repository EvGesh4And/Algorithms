package main

import "fmt"

func main() {
	// Example usage of Merge Sort
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	MergeSort(arr)
	fmt.Println("Sorted array:", arr)
}

func MergeSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	mid := n / 2
	MergeSort(arr[:mid])
	MergeSort(arr[mid:])
	merge(arr, mid)
}

func merge(arr []int, mid int) {
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
		if i < len(left) {
			copy(arr[k:], left[i:])
		}
		if j < len(right) {
			copy(arr[k:], right[j:])
		}
	}
}
