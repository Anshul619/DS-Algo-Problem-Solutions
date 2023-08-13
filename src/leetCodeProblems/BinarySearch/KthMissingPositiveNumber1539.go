package main

/*
- LeetCode - https://leetcode.com/problems/kth-missing-positive-number/description/
- Time - O(logn)
- Space - O(1)
*/

func bsUtil(start, end, k int, arr []int) int {

	if start < 0 || end >= len(arr) || start > end {
		return -1
	}

	mid := start + (end-start)/2

	if mid == start {

		diff := arr[end] - arr[start] - (end - start)

		if diff < k {
			return arr[end] + (k - diff)
		}

		return arr[mid] + k
	}

	leftDiff := arr[mid] - arr[start] - (mid - start)

	if leftDiff < k {
		k -= leftDiff
		return bsUtil(mid, end, k, arr)
	} else {
		return bsUtil(start, mid, k, arr)
	}
}
func findKthPositive(arr []int, k int) int {

	if len(arr) == 0 {
		return k
	}

	diff := arr[0] - 1

	if diff >= k {
		return k
	} else {
		k -= arr[0] - 1
	}

	return bsUtil(0, len(arr)-1, k, arr)
}

// func main() {
// 	//log.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
// 	log.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
// 	//log.Println(findKthPositive([]int{2}, 1))
// 	//log.Println(findKthPositive([]int{7, 13, 21, 25, 29, 32, 38, 45}, 4))
// }
