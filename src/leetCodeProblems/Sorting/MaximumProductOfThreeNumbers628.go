package main

/*
- LeetCode - https://leetcode.com/problems/maximum-product-of-three-numbers/description/
- Time - O(nlogn)
- Space - O(1)
*/
import (
	"log"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	lastThreeProduct := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]
	firstTwoAndLastProduct := nums[0] * nums[1] * nums[len(nums)-1]

	if firstTwoAndLastProduct > lastThreeProduct {
		return firstTwoAndLastProduct
	}

	return lastThreeProduct
}

func main() {
	log.Println(maximumProduct([]int{-1, -2, -3}))
	log.Println(maximumProduct([]int{1, 2, 3, 4}))
	log.Println(maximumProduct([]int{1, 2, 3}))
}
