package main

/*
- LeetCode - https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/submissions/
*/

import (
	"sort"
)

func findMinMax(node *TreeNode, x int, minX *int, maxX *int, y int, minY *int, maxY *int, m map[int]map[int][]int) {

	if node == nil {
		return
	}

	if *minX > x {
		*minX = x
	}

	if *maxX < x {
		*maxX = x
	}

	if *minY > y {
		*minY = y
	}

	if *maxY < y {
		*maxY = y
	}

	if _, ok := m[y]; !ok {
		m[y] = make(map[int][]int)
	}

	m[y][x] = append(m[y][x], node.Val)
	findMinMax(node.Left, x+1, minX, maxX, y-1, minY, maxY, m)
	findMinMax(node.Right, x+1, minX, maxX, y+1, minY, maxY, m)
}

func verticalTraversal(root *TreeNode) [][]int {
	minX, maxX, minY, maxY := 0, 0, 0, 0

	m := make(map[int]map[int][]int)
	findMinMax(root, 0, &minX, &maxX, 0, &minY, &maxY, m)
	// log.Println(m)
	// log.Println(minX, maxX, minY, maxY)
	out := make([][]int, 0)

	for j := minY; j <= maxY; j++ {

		temp := []int{}

		for i := minX; i <= maxX; i++ {

			if len(m[j][i]) > 0 {
				sort.Ints(m[j][i])
			}

			// log.Println(m[j][i])
			temp = append(temp, m[j][i]...)
		}

		out = append(out, temp)
	}
	return out

}

// func main() {
// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 3
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 9
// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 20
// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Left.Val = 15
// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 7

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 1
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 4
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 5

// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 3
// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Left.Val = 6
// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 7

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 0
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 1
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 2
// 	root.Left.Right.Left = new(TreeNode)
// 	root.Left.Right.Left.Val = 6
// 	root.Left.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Val = 3
// 	root.Left.Right.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Right.Val = 4
// 	root.Left.Right.Right.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Right.Right.Val = 5

// 	log.Println(verticalTraversal(root))
// }
