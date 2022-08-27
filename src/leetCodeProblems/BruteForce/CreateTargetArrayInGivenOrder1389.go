package main

/*
- LeetCode - https://leetcode.com/problems/create-target-array-in-the-given-order/
*/
//import "log"

func createTargetArrayBruteForce(nums []int, index []int) []int {

	output := make([]int, len(nums))

	for i := range output {
		output[i] = -1
	}

	for i, v := range index {

		if output[v] != -1 {
			for j := len(output) - 1; j > v; j-- {
				output[j] = output[j-1]
			}
		}

		output[v] = nums[i]
	}

	return output
}

func createTargetArray(nums []int, index []int) []int {

	output := []int{nums[0]}

	for i := 1; i < len(index); i++ {

		if index[i] > len(output)-1 {
			output = append(output, nums[i])
		} else {
			postElements := append([]int{nums[i]}, output[index[i]:]...)
			output = append(output[:index[i]], postElements...)
		}
	}

	return output
}

// func main() {

// 	//nums := []int{0, 1, 2, 3, 4}
// 	//index := []int{0, 1, 2, 2, 1}

// 	nums := []int{1, 2, 3, 4, 0}
// 	index := []int{0, 1, 2, 3, 0}

// 	log.Println(createTargetArray(nums, index))
// }
