package main

/*
- Leetcode - https://leetcode.com/problems/rotate-array/description/
- Space - O(1)
- Time - O(n)
*/

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Rotate array by left k
func leftRotate(nums []int, k int) {
	gcd := gcd(k, len(nums))

	for i := 0; i < gcd; i++ {
		temp := nums[i]
		j := i

		// update values
		for {
			k1 := j + k
			if k1 >= len(nums) {
				k1 = k1 - len(nums)
			}

			if k1 == i {
				break
			}

			nums[j] = nums[k1]
			j = k1
		}

		// update last element of ith set to temp
		nums[j] = temp
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	k = len(nums) - k

	leftRotate(nums, k)
}
