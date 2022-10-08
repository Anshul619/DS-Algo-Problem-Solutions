package main

/*
- LeetCode - https://leetcode.com/problems/maximum-binary-tree/discuss/2674931/Go-Recursion-or-DFS-or-2-Pointers-or-85-Faster-or-85-Less-Memory
*/
import "log"

func maximumBTUtil(nums []int, start int, end int, isLeftChild bool, parent *TreeNode) *TreeNode {

	if start < 0 || end >= len(nums) || start > end {
		return parent
	}

	max := 0
	maxIndex := start

	for i := start; i <= end; i++ {
		if nums[i] > max {
			maxIndex = i
			max = nums[i]
		}
	}

	node := new(TreeNode)
	node.Val = max

	if parent == nil {
		parent = node
	} else {
		if isLeftChild {
			parent.Left = node
		} else {
			parent.Right = node
		}
	}

	maximumBTUtil(nums, start, maxIndex-1, true, node)
	maximumBTUtil(nums, maxIndex+1, end, false, node)

	// log.Println(parent)
	return parent
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return maximumBTUtil(nums, 0, len(nums)-1, true, nil)
}

func printInOrder(node *TreeNode) {

	if node == nil {
		return
	}

	log.Println(node.Val)
	printInOrder(node.Left)

	printInOrder(node.Right)
}

func main() {
	//nums := []int{3, 2, 1, 6, 0, 5}
	nums := []int{3, 2, 1}
	node := constructMaximumBinaryTree(nums)

	printInOrder(node)
}
