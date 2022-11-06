package main

/*
- Leetcode - https://leetcode.com/problems/single-number/
- TimeComplexity - O(n)
- Space Complexity - O(1)
*/
import "log"

func singleNumber(nums []int) int {
	ansXOR := 0

	for _, v := range nums {
		ansXOR ^= v
	}

	return ansXOR

}

func main() {
	nums := []int{4, 1, 2, 1, 2}

	log.Println(singleNumber(nums))
}
