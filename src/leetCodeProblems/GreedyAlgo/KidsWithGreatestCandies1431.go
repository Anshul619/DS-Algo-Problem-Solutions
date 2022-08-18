package main

/*
LeetCode - https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
*/
import "log"

func kidsWithCandies(candies []int, extraCandies int) []bool {

	max := 0
	output := make([]bool, len(candies))

	for _, v := range candies {
		if max < v {
			max = v
		}
	}

	for i, v := range candies {

		if v+extraCandies >= max {
			output[i] = true
		}
	}

	return output
}

func main() {

	//input := []int{2, 3, 5, 1, 3}
	//extraCandies := 3

	//input := []int{4, 2, 1, 1, 2}
	//extraCandies := 1

	input := []int{12, 1, 12}
	extraCandies := 10

	log.Println(kidsWithCandies(input, extraCandies))
}
