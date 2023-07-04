package main

/*
- LeetCode - https://leetcode.com/problems/rearrange-array-elements-by-sign
- Space - O(1)
- Time - O(n)
*/

func rightRotateArray(nums []int, outOfPlace int, curIndex int) {
	cur := nums[curIndex]

	for i := curIndex; i > outOfPlace; i-- {
		nums[i] = nums[i-1]
	}

	nums[outOfPlace] = cur
}

func rearrangeArray(nums []int) []int {

	outOfPlace := -1

	for i := range nums {

		if outOfPlace >= 0 {

			if (nums[outOfPlace] < 0 && nums[i] >= 0) ||
				(nums[outOfPlace] >= 0 && nums[i] < 0) {
				rightRotateArray(nums, outOfPlace, i)

				if i-outOfPlace >= 2 {
					outOfPlace += 2
				} else {
					outOfPlace = -1
				}
			}
		}

		if outOfPlace == -1 {

			if nums[i] >= 0 && i%2 != 0 {
				outOfPlace = i
			} else if nums[i] < 0 && i%2 == 0 {
				outOfPlace = i
			}
		}
	}

	return nums
}

// func main() {
// 	//log.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
// 	//log.Println(rearrangeArray([]int{-1, 1}))
// 	log.Println(rearrangeArray([]int{28, -41, 22, -8, -37, 46, 35, -9, 18, -6, 19, -26, -37, -10, -9, 15, 14, 31}))
// }
