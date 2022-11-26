package main

/*
- LeetCode - https://leetcode.com/problems/guess-number-higher-or-lower/
- Time - O(logn)
- Space - O(1)
*/

func guess(num int) int {

	g := 1

	if num == g {
		return 0
	} else if num > g {
		return -1
	}
	return 1
}

func guessNumber(n int) int {

	low := 1
	high := n
	mid := 1

	for low <= high {

		mid = low + (high-low)/2

		g := guess(mid)

		// log.Println(g, mid)
		if g == 0 {
			return mid
		} else if g == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return mid
}

// func main() {

// 	//log.Println(guessNumber(10))
// 	//log.Println(guessNumber(1))
// 	log.Println(guessNumber(2))
// }
