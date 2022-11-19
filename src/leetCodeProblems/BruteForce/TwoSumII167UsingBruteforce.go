package main

/*
- LeetCode - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
- TimeComplexity - O(n2)
- Space - O(1)
*/

func twoSum(numbers []int, target int) []int {

	out := []int{}

	for i, v := range numbers {

		for j := i + 1; j < len(numbers); j++ {

			if v+numbers[j] == target {
				out = append(out, i+1)
				out = append(out, j+1)
				break
			}
		}
	}

	return out
}

// func main() {
// 	// numbers := []int{2, 7, 11, 15}
// 	// target := 9

// 	// numbers := []int{2, 3, 4}
// 	// target := 6

// 	numbers := []int{-1, 0}
// 	target := -1

// 	log.Println(twoSum(numbers, target))

// }
