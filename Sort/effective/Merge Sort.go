package main

import "fmt"

func main() {
	// Пример использования Merge Sort
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArr := MergeSort(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}

func MergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
