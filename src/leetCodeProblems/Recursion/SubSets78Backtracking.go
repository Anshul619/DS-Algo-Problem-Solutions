package main

/*
- LeetCode - https://leetcode.com/problems/subsets/solutions/464411/official-solution/
- Time - O(N×2N)
- Space - O(N)
*/
import "log"

func subsets(nums []int) [][]int {

	out := [][]int{}

	for i := 0; i < len(nums); i++ {
		backTrack(nums, i, 0, []int{}, &out)
	}

	return out
}

func backTrack(nums []int, combinationLength int, start int, cur []int, out *[][]int) {

	log.Println(combinationLength, cur, start, *out)

	if combinationLength == start {
		log.Println("Insert", cur, combinationLength, start, *out)
		*out = append(*out, cur)
		log.Println("Insert After", *out)
		return
	}

	for j := start; j < len(nums); j++ {
		cur = append(cur, nums[j])
		log.Println(combinationLength, cur, j, *out)
		backTrack(nums, combinationLength, j+1, cur, out)
		cur = cur[:len(cur)-1]

	}
}

// func main() {
// 	//nums := []int{1, 2, 3}
// 	//nums := []int{0}
// 	nums := []int{9, 0, 3, 5, 7}
// 	log.Println(subsets(nums))
// }
