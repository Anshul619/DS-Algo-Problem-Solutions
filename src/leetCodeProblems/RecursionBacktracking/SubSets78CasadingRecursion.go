package main

/*
- LeetCode - https://leetcode.com/problems/subsets/solutions/464411/official-solution/
- Time - O(N×2N)
- Space - O(N×2N)
*/

func subsetsCasadingRecursion(nums []int) [][]int {

	out := [][]int{}
	out = append(out, []int{})

	subsetUtil(&out, 0, nums)

	return out
}

func subsetUtil(out *[][]int, i int, nums []int) {

	if i >= len(nums) {
		return
	}

	temp := [][]int{}

	for _, v := range *out {
		temp1 := append([]int{}, v...)
		temp1 = append(temp1, nums[i])
		temp = append(temp, temp1)
	}

	*out = append(*out, temp...)
	subsetUtil(out, i+1, nums)
}

// func main() {
// 	//nums := []int{1, 2, 3}
// 	//nums := []int{0}
// 	nums := []int{9, 0, 3, 5, 7}
// 	log.Println(subsetsCasadingRecursion(nums))
// }
