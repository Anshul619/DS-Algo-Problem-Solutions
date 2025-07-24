package main

/*
- Leetcode - https://leetcode.com/problems/rotate-array/description/
- Space - O(1)
- Time - O(n)
*/

// Standard Euclidean algorithm to compute the greatest common divisor (GCD) of two integers a and b
// GCD - the largest positive integer that divides each of the integers without leaving a remainder.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func leftRotate(nums []int, k int) {
	gcd := gcd(k, len(nums)) // example - k=4, len=7, gcd=1

	// Creates gcd number of cycles and moves elements within each cycle.
	for i := range gcd {
		t := nums[i]
		j := i

		// Move each element k steps ahead (left rotation)
		for {
			k1 := j + k
			if k1 >= len(nums) {
				k1 = k1 - len(nums) // wrap around
			}

			if k1 == i {
				break
			}

			nums[j] = nums[k1]
			j = k1
		}

		nums[j] = t
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums) // reduce unnecessary full rotations
	k = len(nums) - k // convert right rotation to equivalent left rotation
	leftRotate(nums, k)
}
