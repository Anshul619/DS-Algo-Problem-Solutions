package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	unique := 0

	for i := 1; i < len(nums); i++ {

		// if numbers are same, increase i counter
		// otherwise, swap elements
		if nums[unique] != nums[i] {
			unique++

			// swap elements
			if i != unique {
				t := nums[i]
				nums[i] = nums[unique]
				nums[unique] = t
			}
		}
	}
	return unique + 1
}
