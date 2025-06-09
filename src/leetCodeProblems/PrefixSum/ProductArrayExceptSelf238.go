package main

/*
- Leetcode - https://leetcode.com/problems/product-of-array-except-self/description/
- Time - O(n)
- Space - O(1)
*/
func productExceptSelf(nums []int) []int {

	product := make([]int, len(nums))
	temp := 1

	for i, v := range nums {
		product[i] = temp
		temp = temp * v
	}

	temp = 1

	for i := len(nums) - 1; i >= 0; i-- {
		product[i] = product[i] * temp
		temp = temp * nums[i]
	}

	return product
}
