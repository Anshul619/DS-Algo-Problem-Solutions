package main

/*
- LeetCode - https://leetcode.com/submissions/detail/899901477/
- Time - O(n)
- Space - O(1)
*/

func getMax(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func minimumRefill(plants []int, capacityA int, capacityB int) int {

	out := 0
	currentA, currentB := 0, len(plants)-1
	bucketA, bucketB := capacityA, capacityB

	for currentA <= currentB {

		if currentA == currentB {
			if getMax(bucketA, bucketB) < plants[currentA] {
				out++
			}
			break
		} else {

			if bucketA < plants[currentA] {
				out++
				bucketA = capacityA
			}

			if bucketB < plants[currentB] {
				out++
				bucketB = capacityB
			}

			bucketA -= plants[currentA]
			bucketB -= plants[currentB]

			currentA++
			currentB--

		}
	}

	return out
}

// func main() {
// 	//plants := []int{2, 2, 3, 3}

// 	//log.Println(minimumRefill(plants, 5, 5))
// 	//log.Println(minimumRefill(plants, 3, 4))

// 	plants := []int{5}
// 	log.Println(minimumRefill(plants, 10, 8))
// }
