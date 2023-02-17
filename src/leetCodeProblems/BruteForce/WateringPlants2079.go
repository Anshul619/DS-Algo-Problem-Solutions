package main

/*
- LeetCode - https://leetcode.com/problems/watering-plants/description/
- Space - O(1)
- Time - O(n)
*/
import "log"

func wateringPlants(plants []int, capacity int) int {
	out := 0
	bucket := capacity

	for i, v := range plants {

		if bucket < v {
			out += 2 * i
			bucket = capacity
		}

		bucket -= v
		out++
	}

	return out
}

func main() {
	// plants := []int{2, 2, 3, 3}
	// capacity := 5

	plants := []int{1, 1, 1, 4, 2, 3}
	capacity := 4

	log.Println(wateringPlants(plants, capacity))
}
