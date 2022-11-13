package main

/*
- LeetCode - https://leetcode.com/problems/flood-fill/
- Space - O(n), due to call stack
- Time - O(n)
*/

import "log"

func floodFillUtil(image [][]int, sr int, sc int, targetColor int, startingPixelColor int) {

	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != startingPixelColor || image[sr][sc] == targetColor {
		return
	}

	image[sr][sc] = targetColor

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	for i := 0; i < 4; i++ {

		x := sr + dx[i]
		y := sc + dy[i]

		floodFillUtil(image, x, y, targetColor, startingPixelColor)

	}
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	floodFillUtil(image, sr, sc, color, image[sr][sc])
	return image

}

func main() {
	// image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	// sr := 1
	// sc := 1
	// color := 2

	image := [][]int{{0, 0, 0}, {0, 0, 0}}
	sr := 0
	sc := 0
	color := 0

	log.Println(floodFill(image, sr, sc, color))
}
