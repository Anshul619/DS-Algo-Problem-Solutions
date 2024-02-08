package main

/*
- LeetCode - https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
- Time - O(n)
- Space - O(1)
*/

func findDuplicates(nums []int) []int {
	n := len(nums) + 1

	out := []int{}

	for _, v := range nums {
		base := v % n
		nums[base-1] += n
	}

	for i, v := range nums {
		if v/n == 2 {
			out = append(out, i+1)
		}
	}
	return out
}

// func main() {
// 	log.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
// 	log.Println(findDuplicates([]int{1, 1, 2}))
// 	log.Println(findDuplicates([]int{1}))
// 	log.Println(findDuplicates([]int{10, 2, 5, 10, 9, 1, 1, 4, 3, 7}))
// }
