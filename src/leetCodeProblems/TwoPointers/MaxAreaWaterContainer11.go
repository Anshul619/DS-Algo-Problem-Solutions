package main

/*
- Leetcode - https://leetcode.com/problems/container-with-most-water/description/
- Space - O(1)
- Time - O(n)
*/

func minH(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func maxArea(height []int) int {

	max := 0
	first := 0
	last := len(height) - 1

	for first < last {
		temp := (last - first) * minH(height[first], height[last])

		if temp > max {
			max = temp
		}

		if height[first] < height[last] {
			first++
		} else {
			last--
		}
	}

	return max
}

// func main() {
// 	log.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// 	log.Println(maxArea([]int{1, 1}))
// }
