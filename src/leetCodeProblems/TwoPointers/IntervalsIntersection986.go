package main

/*
- Leetcode - https://leetcode.com/problems/interval-list-intersections/description/
- Time - O(m+n)
- Space - O(1)
*/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var out [][]int
	first := 0
	second := 0

	for first < len(firstList) && second < len(secondList) {

		lower := max(firstList[first][0], secondList[second][0])
		high := min(firstList[first][1], secondList[second][1])

		if lower <= high {
			out = append(out, []int{lower, high})
		}

		if firstList[first][1] < secondList[second][1] {
			first++
		} else {
			second++
		}
	}
	return out
}

// func main() {
// 	log.Println(intervalIntersection([][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}, [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}))
// }
