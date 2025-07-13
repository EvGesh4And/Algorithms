package main

import "fmt"

func main() {
	// Example usage
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := findKthLargest(nums, k)
	fmt.Println("The", k, "th largest element is:", result)
}
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	heap := nums[:k]

	for i := k/2 - 1; i >= 0; i-- {
		heapify(nums, k, i)
	}

	for i := k; i < n; i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			heapify(heap, k, 0)
		}
	}
	return heap[0]
}

func heapify(arr []int, n, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < n && arr[right] < arr[smallest] {
		smallest = right
	}
	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		heapify(arr, n, smallest)
	}
}
