package main

/*
- LeetCode - https://leetcode.com/problems/rearrange-array-elements-by-sign
- Space - O(n)
- Time - O(n)
*/

// func rearrangeArray(nums []int) []int {

// 	allNegative := []int{}
// 	allPositive := []int{}

// 	for _, v := range nums {
// 		if v > 0 {
// 			allPositive = append(allPositive, v)
// 		} else {
// 			allNegative = append(allNegative, v)
// 		}
// 	}

// 	for i := 0; i < len(nums)/2; i++ {
// 		nums[2*i] = allPositive[i]
// 		nums[2*i+1] = allNegative[i]
// 	}

// 	return nums
// }

func rearrangeArrayUsingSpace(nums []int) []int {

	positiveIndex, negativeIndex := 0, 1
	out := make([]int, len(nums))

	for _, v := range nums {
		if v > 0 {
			out[positiveIndex] = v
			positiveIndex = positiveIndex + 2
		} else {
			out[negativeIndex] = v
			negativeIndex = negativeIndex + 2
		}
	}

	return out
}

// func main() {
// 	//log.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
// 	//log.Println(rearrangeArray([]int{-1, 1}))
// 	log.Println(rearrangeArrayUsingSpace([]int{28, -41, 22, -8, -37, 46, 35, -9, 18, -6, 19, -26, -37, -10, -9, 15, 14, 31}))
// }
