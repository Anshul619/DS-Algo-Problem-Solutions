package main

/*
- LeetCode - https://leetcode.com/problems/unique-number-of-occurrences/
- Time - O(n)
- Space - O(n)
*/

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)

	for _, v := range arr {
		m[v]++
	}

	occurences := make(map[int]struct{})

	for _, v := range m {
		if _, ok := occurences[v]; ok {
			return false
		}
		occurences[v] = struct{}{}
	}

	return true
}

// func main() {
// 	log.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
// 	log.Println(uniqueOccurrences([]int{1, 2}))
// 	log.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
// }
