package main

/*
- LeetCode - https://leetcode.com/problems/flood-fill/
- Space - O(m*n)
- Time - O(m*n)
*/

var dx = []int{1, 0, -1, 0}
var dy = []int{0, -1, 0, 1}

func floodFillDFS(image [][]int, sr int, sc int, newColor int, originalColor int) {
	if sr < 0 || sr >= len(image) {
		return
	}

	if sc < 0 || sc >= len(image[0]) {
		return
	}

	if image[sr][sc] != originalColor {
		return
	}

	image[sr][sc] = newColor

	for i := range 4 {
		//slog.Info("loop", "i", i, "sr", sr+dx[i], "sc", sc+dy[i])
		floodFillDFS(image, sr+dx[i], sc+dy[i], newColor, originalColor)
	}
}
func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	originalColor := image[sr][sc]
	if image[sr][sc] == color {
		return image
	}

	floodFillDFS(image, sr, sc, color, originalColor)
	return image
}
