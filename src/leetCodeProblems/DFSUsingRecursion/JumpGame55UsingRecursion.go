package main

/*
- LeetCode - https://leetcode.com/problems/jump-game/

Fix needed
- With the example long UC, it goes infinite loop. This needs fix.
*/
import "log"

func canJumpUtil(curIndex int, nums []int) bool {

	if curIndex >= len(nums)-1 {
		return true
	}

	for i := curIndex + 1; i <= curIndex+nums[curIndex]; i++ {
		if canJumpUtil(i, nums) {
			return true
		}
	}

	return false
}

func canJump(nums []int) bool {
	return canJumpUtil(0, nums)
}

func main() {
	//nums := []int{2, 3, 1, 1, 4}

	nums := []int{3, 2, 1, 0, 4}
	log.Println(canJump(nums))
}
