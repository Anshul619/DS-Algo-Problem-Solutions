package main

/*
	- LeetCode - https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
*/
import (
	"log"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {

	sort.Ints(seats)
	sort.Ints(students)

	output := 0

	for i, v := range students {

		if v-seats[i] < 0 {
			output += -(v - seats[i])
		} else {
			output += v - seats[i]
		}
	}

	return output

}

func main() {

	//seats := []int{3, 1, 5}
	//students := []int{2, 7, 4}

	seats := []int{4, 1, 5, 9}
	students := []int{1, 3, 2, 6}

	log.Println(minMovesToSeat(seats, students))
}
