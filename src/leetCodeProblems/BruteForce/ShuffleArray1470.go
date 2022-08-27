package main

//import "log"

func shuffle(nums []int, n int) []int {

	output := make([]int, n*2)

	for i, j := 0, 0; i < len(output); i, j = i+2, j+1 {
		output[i] = nums[j]
		output[i+1] = nums[j+n]
	}

	return output
}

// func main() {

// 	input := []int{2, 5, 1, 3, 4, 7}
// 	n := 3

// 	log.Println(shuffle(input, n))
// }
