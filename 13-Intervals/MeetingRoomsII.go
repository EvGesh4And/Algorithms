package main

import "sort"

type Interval struct {
	Start, End int
}

func MinMeetingRooms(intervals []*Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start <= intervals[j].Start
	})
	minHeap := make([]int, 0)
	for _, interval := range intervals {
		n := len(minHeap)
		if n == 0 || minHeap[0] > interval.Start {
			minHeap = append(minHeap, interval.End)
			shiftUp(minHeap, n)
		} else {
			minHeap[0] = interval.End
			heapify(minHeap, n, 0)
		}
	}

	return len(minHeap)
}

func heapify(arr []int, n, i int) {
	idx := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] < arr[idx] {
		idx = left
	}
	if right < n && arr[right] < arr[idx] {
		idx = right
	}
	if idx != i {
		arr[idx], arr[i] = arr[i], arr[idx]
		heapify(arr, n, idx)
	}
}

func shiftUp(arr []int, i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if arr[i] < arr[parent] {
			arr[i], arr[parent] = arr[parent], arr[i]
			i = parent
		} else {
			break
		}
	}
}
