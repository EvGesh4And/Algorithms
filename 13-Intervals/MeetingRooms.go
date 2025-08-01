package main

import (
	"sort"
)

// LeetCode

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := range len(intervals) - 1 {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}

// LintCode

type Interval struct {
	Start, End int
}

// import "sort"

func CanAttendMeetings(intervals []*Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].End > intervals[i+1].Start {
			return false
		}
	}
	return true
}
