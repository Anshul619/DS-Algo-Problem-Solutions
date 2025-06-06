package main

/*
- Leetcode - https://leetcode.com/problems/insert-interval/description/
- Time - O(n)
- Space - O(n)
*/

func insert1(intervals [][]int, newInterval []int) [][]int {
	out := [][]int{}

	for _, v := range intervals {

		if newInterval == nil {
			out = append(out, v)
			continue
		}

		if v[0] <= newInterval[0] && newInterval[0] <= v[1] {
			if newInterval[1] < v[1] {
				newInterval = []int{v[0], v[1]}
			} else {
				newInterval = []int{v[0], newInterval[1]}
			}
		} else if newInterval[0] <= v[0] && v[0] <= newInterval[1] {
			if v[1] > newInterval[1] {
				newInterval = []int{newInterval[0], v[1]}
			}
		} else { // v and newInterval are non-overlapping

			// there won't be any overlap in future, hence insert in out
			if newInterval[0] < v[0] {
				out = append(out, newInterval)
				newInterval = nil
			}

			out = append(out, v)
		}
	}

	if newInterval != nil {
		out = append(out, newInterval)
	}
	return out
}
