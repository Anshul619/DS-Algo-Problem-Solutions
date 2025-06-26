package main

func heapAlgo(nums []int, size int, out *[][]int) {

	if size == 0 {
		copy := append([]int(nil), nums...)
		*out = append(*out, copy)
		return
	}

	for i := 0; i < size; i++ {
		heapAlgo(nums, size-1, out)

		if size%2 == 0 {
			nums[i], nums[size-1] = nums[size-1], nums[i]
		} else {
			nums[0], nums[size-1] = nums[size-1], nums[0]
		}
	}
}

func permuteHeapAlgo(nums []int) [][]int {

	out := [][]int{}

	heapAlgo(nums, len(nums), &out)

	return out
}

// func main() {

// 	//nums := []int{1, 2, 3}
// 	nums := []int{1, 2, 3, 4}
// 	//nums := []int{1, 2, 3}
// 	//nums := []int{0, 1}
// 	//nums := []int{1}
// 	//nums := []int{}
// 	log.Println(permuteHeapAlgo(nums))
// }
