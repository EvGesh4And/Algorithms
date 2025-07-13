package main

import "fmt"

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	result := lastStoneWeight(stones)
	fmt.Println(result) // Output: 1
}

func lastStoneWeight(stones []int) int {
	n := len(stones)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(stones, n, i)
	}

	for len(stones) > 1 {
		y := stones[0]
		stones[0] = stones[n-1]
		n--
		stones = stones[:n]
		heapify(stones, n, 0)
		x := stones[0]

		if y == x {
			stones[0] = stones[n-1]
			n--
			stones = stones[:n]
			heapify(stones, n, 0)
		} else {
			stones[0] = y - x
			stones = stones[:n]
			heapify(stones, n, 0)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	} else {
		return 0
	}
}

func heapify(arr []int, n, i int) {
	idxMax := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[idxMax] {
		idxMax = left
	}
	if right < n && arr[right] > arr[idxMax] {
		idxMax = right
	}
	if i != idxMax {
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
		heapify(arr, n, idxMax)
	}
}
