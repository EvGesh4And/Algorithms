package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	var res [][]int
	for _, interval := range intervals {
		n := len(res)
		if n == 0 || res[n-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[n-1][1] = max(res[n-1][1], interval[1])
		}
	}
	return res
}

// LintCode

// import "sort"

func Merge(intervals []*Interval) []*Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var res []*Interval

	for _, interval := range intervals {
		n := len(res)
		if n == 0 || res[n-1].End < interval.Start {
			res = append(res, interval)
		} else {
			res[n-1].End = max(res[n-1].End, interval.End)
		}
	}
	return res
}

func max(a, b int) (c int) {
	c = a
	if b > a {
		c = b
	}
	return
}
