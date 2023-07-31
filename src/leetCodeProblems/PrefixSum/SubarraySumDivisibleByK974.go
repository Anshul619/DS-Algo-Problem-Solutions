package main

/*
- LeetCode - https://leetcode.com/problems/subarray-sums-divisible-by-k/
- Time - O(n)
- Extra Space - O(k)
- Submission - https://leetcode.com/problems/subarray-sums-divisible-by-k/solutions/3830508/go-prefix-sum-hashmap-100-faster-100-less-space/
*/
func subarraysDivByK(nums []int, k int) int {
	m := make([]int, k)
	m[0] = 0

	out, sum := 0, 0

	for _, v := range nums {
		sum += v

		// since sum can be negative, hence double modulo
		mod := ((sum % k) + k) % k

		out += m[mod]
		m[mod]++
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
