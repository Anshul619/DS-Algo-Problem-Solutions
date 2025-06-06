package main

/*
- Leetcode - https://leetcode.com/problems/insert-interval/description/
- Time - O(nlogn)
- Space - O(1)
*/
import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {

	intervals = append(intervals, newInterval)

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	lastIndex := 0

	for i := 1; i < len(intervals); i++ {
		// overlapping intervals
		if intervals[i][0] <= intervals[lastIndex][1] {
			intervals[lastIndex][1] = max(intervals[i][1], intervals[lastIndex][1])
		} else {
			lastIndex++
			intervals[lastIndex] = intervals[i]
		}
	}

	return intervals[:lastIndex+1]
}
