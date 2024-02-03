package main

/*
- LeetCode - https://leetcode.com/problems/contains-duplicate/
- Time - O(n)
- Space - O(n)
*/

func containsDuplicate(nums []int) bool {

	m := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}

// func main() {
// 	log.Println(twoSum([]int{2, 7, 11, 15}, 9))
// 	log.Println(twoSum([]int{3, 2, 4}, 6))
// 	log.Println(twoSum([]int{3, 3}, 6))
// }
