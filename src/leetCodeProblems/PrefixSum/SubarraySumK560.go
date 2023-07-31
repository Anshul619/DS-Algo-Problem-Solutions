package main

/*
- LeetCode - https://leetcode.com/problems/subarray-sum-equals-k/
- Time - O(n)
- Extra Space - O(n)
*/

func subarraySum(nums []int, k int) int {

	m := make(map[int]int)
	m[0] = 1

	out, sum := 0, 0

	for _, v := range nums {
		sum += v

		if v, ok := m[sum-k]; ok {
			out += v
		}
		m[sum]++
	}

	return out
}

// func main() {
// 	log.Println(subarraySum([]int{1, 1, 1}, 2))
// 	log.Println(subarraySum([]int{1, 2, 3}, 3))
// 	log.Println(subarraySum([]int{1}, 0))
// }
