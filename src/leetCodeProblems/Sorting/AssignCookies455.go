package main

/*
- LeetCode - https://leetcode.com/problems/assign-cookies/submissions/
- Time - O(nlogn)
- Space - O(1)
*/
import (
	"log"
	"sort"
)

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	//ans := 0
	j := 0

	for _, v := range s {

		if j >= len(g) {
			break
		}

		if v >= g[j] {
			j++
		}

		// ans++
	}

	return j
}

func main() {

	// g := []int{1, 2, 3}
	// s := []int{1, 1}

	// log.Println(findContentChildren(g, s))

	// g := []int{1, 1}
	// s := []int{1, 2, 3}

	// log.Println(findContentChildren(g, s))

	// g := []int{1, 2, 3}
	// s := []int{}

	// log.Println(findContentChildren(g, s))

	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}

	log.Println(findContentChildren(g, s))
}
