package main

/*
- LeetCode - https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/description/
- Space - O(n)
- Time - O(n)
*/

func findThePrefixCommonArray(A []int, B []int) []int {

	mA := make(map[int]bool)
	mB := make(map[int]bool)

	out := make([]int, len(A))

	for i, v := range A {

		if i > 0 {
			out[i] = out[i-1]
		}

		if v == B[i] {
			out[i]++
		} else {
			if _, ok := mB[v]; ok {
				out[i]++
			}

			if _, ok := mA[B[i]]; ok {
				out[i]++
			}
		}

		mA[v] = true
		mB[B[i]] = true
	}

	return out
}

// func main() {
// 	log.Println(findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
// 	log.Println(findThePrefixCommonArray([]int{2, 3, 1}, []int{3, 1, 2}))
// 	log.Println(findThePrefixCommonArray([]int{1, 2}, []int{1, 2}))
// }
