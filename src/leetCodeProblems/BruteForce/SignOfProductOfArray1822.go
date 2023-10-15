package main

/*
- Leetcode - https://leetcode.com/problems/sign-of-the-product-of-an-array/description/
- Time - O(n)
- Space - O(1)
*/
import "log"

func arraySign(nums []int) int {
	n := 0

	for _, v := range nums {
		if v == 0 {
			return 0
		}

		if v < 0 {
			n++
		}
	}

	if n%2 == 1 {
		return -1
	}

	return 1
}

func main() {
	log.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	log.Println(arraySign([]int{1, 5, 0, 2, -3}))
	log.Println(arraySign([]int{-1, 1, -1, 1, -1}))
}
