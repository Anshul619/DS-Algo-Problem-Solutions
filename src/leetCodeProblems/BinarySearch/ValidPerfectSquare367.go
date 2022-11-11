package main

/*
- LeetCode - https://leetcode.com/problems/valid-perfect-square/
- Time - O(logn)
- Space - O(1)
*/
import "log"

// func isPerfectSquare(num int) bool {

// 	if num == 1 {
// 		return true
// 	}

// 	low := 1
// 	high := num

// 	for low < high {

// 		mid := (low + high) / 2

// 		// log.Println(low, mid, high)
// 		currentSquare := mid * mid

// 		if currentSquare == num {
// 			return true
// 		} else if currentSquare > num {
// 			high = mid - 1
// 		} else {
// 			low = mid + 1
// 		}
// 	}

// 	return false
// }

func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}

	low := 1
	high := num

	for low < high {

		mid := low + (high-low)/2

		log.Println(low, mid, high)
		currentSquare := mid * mid

		if currentSquare == num {
			return true
		} else if currentSquare > num {

			if high == mid {
				break
			}

			high = mid
		} else {

			if low == mid {
				break
			}

			low = mid
		}
	}

	return false
}

// func main() {
// 	//num := 16

// 	// num := 14

// 	num := 104976

// 	log.Println(isPerfectSquare(num))
// }
