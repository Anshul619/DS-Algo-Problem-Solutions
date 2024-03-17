package main

/*
- Leetcode - https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
- Time - O(nlogn)
- Space - O(1)
*/
import (
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	out := 0

	overlappingEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] <= overlappingEnd {
			overlappingEnd = min(overlappingEnd, points[i][1])
		} else {
			out++
			overlappingEnd = points[i][1]
		}
	}
	return out + 1
}
