package main

/*
- Leetcode - https://leetcode.com/problems/average-of-levels-in-binary-tree
- Time - O(h*w)
- Space - O(w)
*/

func averageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return []float64{}
	}

	q := new(nodesQueue)
	out := []float64{}

	q.push(root)

	for !q.isEmpty() {

		qt := new(nodesQueue)
		sum, count := 0, 0

		// traverse current level
		for !q.isEmpty() {
			node := q.pop()

			sum += node.Val
			count++

			if node.Left != nil {
				qt.push(node.Left)
			}

			if node.Right != nil {
				qt.push(node.Right)
			}
		}

		q = qt
		out = append(out, float64(sum)/float64(count))

	}
	return out
}
