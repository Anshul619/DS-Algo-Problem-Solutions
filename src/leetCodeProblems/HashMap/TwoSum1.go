package main

/*
- LeetCode - https://leetcode.com/problems/two-sum/solutions/3723755/go-hash-map-100-faster/
- Time - O(n)
- Space - O(n)
*/
import "log"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}

	return []int{}
}

func main() {
	log.Println(twoSum([]int{2, 7, 11, 15}, 9))
	log.Println(twoSum([]int{3, 2, 4}, 6))
	log.Println(twoSum([]int{3, 3}, 6))
}