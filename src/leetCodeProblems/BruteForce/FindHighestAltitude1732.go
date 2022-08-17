package main

/*
- LeetCode - https://leetcode.com/problems/find-the-highest-altitude/submissions/
*/
import (
	"log"
)

func largestAltitude(gain []int) int {

	current, max := 0, 0

	for _, value := range gain {
		current += value

		if current > max {
			max = current
		}
	}

	return max
}

func main() {

	input := []int{-5, 1, 5, 0, -7}
	log.Println(largestAltitude(input))
}
