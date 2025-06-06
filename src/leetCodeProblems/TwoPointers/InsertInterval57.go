package main

import "math"

/*
- Leetcode - https://leetcode.com/problems/insert-interval/description/
- Time - O(n)
- Space - O(n)
*/

func getInsertIndex(intervals [][]int, newInterval []int) int {

	if intervals[0][0] >= newInterval[0] {
		return 0
	}

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][0] <= newInterval[0] && intervals[i+1][0] >= newInterval[0] {
			return i + 1
		}
	}

	return len(intervals)
}

func mergeIntervals2(intervals [][]int) [][]int {
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

func insert2(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	insertIndex := getInsertIndex(intervals, newInterval)

	newIntervals := [][]int{}
	newIntervals = append(newIntervals, intervals[:insertIndex]...)
	newIntervals = append(newIntervals, newInterval)
	newIntervals = append(newIntervals, intervals[insertIndex:]...)

	return mergeIntervals(newIntervals)
}
