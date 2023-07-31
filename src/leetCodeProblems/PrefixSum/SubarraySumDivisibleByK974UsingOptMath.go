package main

/*
- LeetCode - https://leetcode.com/problems/subarray-sums-divisible-by-k/
- Time - O(n+k)
- Extra Space - O(k)
- Submission - https://leetcode.com/problems/subarray-sums-divisible-by-k/solutions/3830469/go-prefix-sum-100-faster-90-less-space/
*/
func subarraysDivByKUsingOptimizedMath(nums []int, k int) int {
	m := make([]int, k)
	m[0] = 0

	out, sum := 0, 0

	for _, v := range nums {
		sum += v
		m[((sum%k)+k)%k]++
	}

	for _, v := range m {
		if v > 1 {
			out += (v * (v - 1)) / 2
		}
	}

	out += m[0]
	return out
}

// func main() {
// 	// log.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
// 	// log.Println(subarraysDivByK([]int{5}, 9))
// 	// log.Println(subarraysDivByK([]int{-1, 2, 9}, 2))
// 	log.Println(subarraysDivByK([]int{2, -2, 2, -4}, 6))
// }
