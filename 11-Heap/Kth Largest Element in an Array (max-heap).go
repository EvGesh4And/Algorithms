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
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	for i := n - 1; i > n-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
	return nums[0]
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
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}
