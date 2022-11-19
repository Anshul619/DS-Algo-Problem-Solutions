package main

/*
- LeetCode - https://leetcode.com/problems/jump-game/
- Time - O(n)
- Space - O(1)
*/
import "log"

func canJump(nums []int) bool {

	max := 0
	newJump := 0

	for i, v := range nums {

		if i > max {
			return false
		}

		newJump = i + v
		if max < newJump {
			max = newJump
		}
	}

	return true
}

func main() {
	nums := []int{2, 3, 1, 1, 4}

	//nums := []int{3, 2, 1, 0, 4}
	log.Println(canJump(nums))
}
