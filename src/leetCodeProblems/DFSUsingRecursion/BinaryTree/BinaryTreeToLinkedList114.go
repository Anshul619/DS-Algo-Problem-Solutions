package main

/*
- Leetcode - https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
- Time - O(n)
- Space - O(1)
*/
func flattenUtil(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	// leaf node
	if node.Left == nil && node.Right == nil {
		return node
	}

	leftEnd := flattenUtil(node.Left)
	rightEnd := flattenUtil(node.Right)

	tempRight := node.Right

	if rightEnd == nil {
		rightEnd = leftEnd
	}

	if leftEnd != nil {
		leftEnd.Right = tempRight
	}

	if node.Left != nil {
		node.Right = node.Left
	}

	node.Left = nil
	return rightEnd
}

func flatten(root *TreeNode) {
	flattenUtil(root)
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 3
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 4

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 5
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 6

// 	// flatten(root)

// 	// printInOrder(root)

// 	root = new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 3
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 4
// 	root.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Val = 5

// 	flatten(root)

// 	printInOrder(root)

// 	//printLinkedList(root)

// 	// head = new(ListNode)
// 	// head.Val = 5
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 6
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 3

// 	// head1 = new(ListNode)
// 	// head1.Val = 8
// 	// head1.Next = new(ListNode)
// 	// head1.Next.Val = 4
// 	// head1.Next.Next = new(ListNode)
// 	// head1.Next.Next.Val = 2

// 	// out = addTwoNumbers(head, head1)

// 	// printLinkedList(out)
// }
