package main

/*
- LeetCode - https://leetcode.com/problems/subsets/solutions/464411/official-solution/
- Time - O(N×2N)
- Space - O(N×2N)
*/
func subsetsCasading(nums []int) [][]int {

	out := [][]int{}
	out = append(out, []int{})

	for _, v := range nums {

		temp := [][]int{}

		for _, s := range out {
			temp1 := append([]int{}, s...)
			temp1 = append(temp1, v)
			temp = append(temp, temp1)
		}

		out = append(out, temp...)
	}

	return out
}

// func main() {
// 	//nums := []int{1, 2, 3}
// 	//nums := []int{0}
// 	nums := []int{9, 0, 3, 5, 7}
// 	log.Println(subsetsCasading(nums))
// }
