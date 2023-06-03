package main

/*
- LeetCode - https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/description/
- Time - O(n)
- Space - O(1)
*/
func bstFromPreorderUtil(preorder []int, start, end int) *TreeNode {

	if start < 0 || end >= len(preorder) || end < start {
		return nil
	}

	node := new(TreeNode)
	node.Val = preorder[start]

	if start == end {
		return node
	}

	rightChildIndex := end + 1

	for i := start; i <= end; i++ {
		if preorder[i] > preorder[start] {
			rightChildIndex = i
			break
		}
	}

	node.Left = bstFromPreorderUtil(preorder, start+1, rightChildIndex-1)
	node.Right = bstFromPreorderUtil(preorder, rightChildIndex, end)

	return node
}

func bstFromPreorder(preorder []int) *TreeNode {
	return bstFromPreorderUtil(preorder, 0, len(preorder)-1)
}

func main() {
	preorder := []int{8, 5, 1, 7, 10, 12}
	//preorder := []int{4, 2}
	printTree(bstFromPreorder(preorder))
}
