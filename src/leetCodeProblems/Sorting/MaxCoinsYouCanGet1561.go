package main

/*
- LeetCode - https://leetcode.com/problems/maximum-number-of-coins-you-can-get/description/
- Time - O(nlogn)
- Space - O(n)
*/
import (
	"log"
	"sort"
)

func maxCoins(piles []int) int {

	sort.Ints(piles)
	out := 0

	first := 0
	last := len(piles) - 2

	for first < last {
		out += piles[last]
		first++
		last -= 2
	}

	return out
}

func main() {

	//piles := []int{2, 4, 1, 2, 7, 8}

	piles := []int{9, 8, 7, 6, 5, 1, 2, 3, 4}
	log.Println(maxCoins(piles))
}
