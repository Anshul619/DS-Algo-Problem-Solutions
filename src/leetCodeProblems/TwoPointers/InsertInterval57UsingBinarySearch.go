package main

import "math"

/*
- Leetcode - https://leetcode.com/problems/insert-interval/description/
- Time - O(n)
- Space - O(n)
*/
func findInsertIndex(intervals [][]int, newInterval []int) int {

	start := 0
	end := len(intervals)

	for start < end {
		mid := start + (end-start)/2

		if intervals[mid][0] < newInterval[0] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}

func mergeIntervals(intervals [][]int) [][]int {
	out := [][]int{}

	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if end >= intervals[i][0] {
			end = int(math.Max(float64(intervals[i][1]), float64(end)))
		} else {
			out = append(out, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}

	return append(out, []int{start, end})
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	insertIndex := findInsertIndex(intervals, newInterval)

	newIntervals := [][]int{}
	newIntervals = append(newIntervals, intervals[:insertIndex]...)
	newIntervals = append(newIntervals, newInterval)
	newIntervals = append(newIntervals, intervals[insertIndex:]...)

	return mergeIntervals(newIntervals)
}
