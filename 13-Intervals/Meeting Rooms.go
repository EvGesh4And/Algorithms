package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canAttendMeetings([][]int{{1, 3}, {4, 9}, {4, 7}}))
}

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	for i := range len(intervals) - 1 {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}
