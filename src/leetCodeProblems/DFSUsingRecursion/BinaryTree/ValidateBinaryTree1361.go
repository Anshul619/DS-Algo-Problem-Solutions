package main

/*
- LeetCode - https://leetcode.com/problems/validate-binary-tree-nodes/
- Time - O(n)
- Space - O(n)
*/

// find cycle through DFS
func hasCycle(node int, visited []bool, leftChild, rightChild []int) bool {
	if node == -1 {
		return false
	}

	if visited[node] {
		return true
	}
	visited[node] = true
	return hasCycle(leftChild[node], visited, leftChild, rightChild) || hasCycle(rightChild[node], visited, leftChild, rightChild)
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	parent := make([]int, n)

	for i := range parent {
		parent[i] = -1
	}

	for i, v := range leftChild {
		if v != -1 {

			if parent[v] != -1 {
				return false
			}

			parent[v] = i
		}
	}

	for i, v := range rightChild {
		if v != -1 {

			if parent[v] != -1 {
				return false
			}

			parent[v] = i
		}
	}

	root := -1

	// find any node without parent
	for i, v := range parent {
		if v == -1 {
			if root != -1 { // more than one root
				return false
			}
			root = i
		}
	}

	// // no root node
	if root == -1 {
		return false
	}

	visited := make([]bool, n)

	isCycleExists := hasCycle(root, visited, leftChild, rightChild)
	if isCycleExists {
		return false
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
