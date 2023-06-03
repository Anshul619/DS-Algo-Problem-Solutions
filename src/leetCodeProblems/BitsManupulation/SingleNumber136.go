package main

/*
- Leetcode - https://leetcode.com/problems/single-number/
- TimeComplexity - O(n)
- Space Complexity - O(1)
*/

func singleNumber1(nums []int) int {
	ansXOR := 0

	for _, v := range nums {
		ansXOR ^= v
	}

	return ansXOR

}

// func main() {
// 	//nums := []int{4, 1, 2, 1, 2}

// 	//nums := []int{0, 1, 0, 1, 0, 1, 99}
// 	nums := []int{2, 2, 3, 2}
// 	log.Println(singleNumber1(nums))
// }
