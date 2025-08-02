package main

import "fmt"

func main() {
	piles := []int{4, 3, 6, 7}
	fmt.Println(minStoneSum(piles, 3))
}

func minStoneSum(piles []int, k int) int {
	n := len(piles)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(piles, n, i)
	}
	for _ = range k {
		piles[0] -= piles[0] / 2
		heapify(piles, n, 0)
	}
	res := 0
	for _, v := range piles {
		res += v
	}
	return res
}

func heapify(arr []int, n, i int) {
	idx := i
	left := 2*i + 1
	right := 2*i + 2
	for left < n && arr[left] > arr[idx] {
		idx = left
	}
	for right < n && arr[right] > arr[idx] {
		idx = right
	}
	if idx != i {
		arr[i], arr[idx] = arr[idx], arr[i]
		heapify(arr, n, idx)
	}
}
