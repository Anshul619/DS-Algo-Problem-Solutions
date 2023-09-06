package main

/*
- Leetcode - https://leetcode.com/problems/can-place-flowers/solutions/4002498/go-simple-greedy-70-faster/
- Time - O(n)
- Space - O(1)
*/

func canPlaceFlowers(flowerbed []int, n int) bool {

	i := 0

	for i < len(flowerbed) {
		if flowerbed[i] == 0 &&
			(i+1 >= len(flowerbed) || flowerbed[i+1] == 0) &&
			(i == 0 || flowerbed[i-1] == 0) {
			n--
			flowerbed[i] = 1
			i = i + 2
		} else {
			i++
		}

		if n == 0 {
			break
		}
	}
	return n == 0
}

// func main() {
// 	// log.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
// 	// log.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
// 	log.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1}, 2))
// 	log.Println(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1))

// }
