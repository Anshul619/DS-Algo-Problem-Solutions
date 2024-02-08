package main

/*
- LeetCode - https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/description/
- Time - O(n)
- Space - O(n)
*/
import "log"

func findSmallestSetOfVertices(n int, edges [][]int) []int {

	degrees := make([]int, n)
	out := []int{}

	for _, v := range edges {
		degrees[v[1]]++
	}

	for i, v := range degrees {
		if v == 0 {
			out = append(out, i)
		}
	}
	return out
}

func main() {

	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}

	log.Println(findSmallestSetOfVertices(n, edges))
}
