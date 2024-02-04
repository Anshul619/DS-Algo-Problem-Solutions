package main

/*
- LeetCode - https://leetcode.com/problems/permutations
*/
func pUtil(input [][]int, num int) [][]int {

	out := [][]int{}

	for _, v := range input {

		for i := 0; i <= len(v); i++ {

			copy := append([]int(nil), v...)

			end := []int{num}
			end = append(end, copy[i:]...)
			out = append(out, append(copy[:i], end...))
		}
	}

	return out
}

func permute(nums []int) [][]int {

	len := len(nums)

	if len == 0 {
		return [][]int{}
	}

	out := [][]int{{nums[len-1]}}
	for i := len - 2; i > -1; i-- {
		out = pUtil(out, nums[i])
	}

	return out
}

// func main() {

// 	//nums := []int{1, 2, 3}
// 	nums := []int{1, 2, 3, 4}
// 	//nums := []int{1, 2, 3}
// 	//nums := []int{0, 1}
// 	//nums := []int{1}
// 	//nums := []int{}
// 	log.Println(permute(nums))
// }
