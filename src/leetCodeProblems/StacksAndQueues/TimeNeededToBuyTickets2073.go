package main

/*
LeetCode - https://leetcode.com/problems/time-needed-to-buy-tickets
*/
//import "log"

func timeRequiredToBuy(tickets []int, k int) int {

	out := 0
	endMainLoop := false

	for !endMainLoop {

		end := len(tickets)

		if tickets[k] == 1 {
			end = k + 1
			endMainLoop = true
		}

		for i := 0; i < end; i++ {

			elem := tickets[0]

			if elem > 0 {
				elem--
				out++
			}

			tickets = append(tickets[1:], elem)
		}
	}

	return out
}

// func main() {

// 	//tickets := []int{2, 3, 2}
// 	//k := 2

// 	//tickets := []int{5, 1, 1, 1}
// 	//k := 0

// 	tickets := []int{84, 49, 5, 24, 70, 77, 87, 8}
// 	k := 3

// 	//tickets := []int{32, 90, 5, 48, 69, 21, 41, 55, 32}
// 	//k := 2

// 	//[]int{1, 2, 1}
// 	//[]int{0, 1, 0} // 6 seconds

// 	log.Println(timeRequiredToBuy(tickets, k))
// }
