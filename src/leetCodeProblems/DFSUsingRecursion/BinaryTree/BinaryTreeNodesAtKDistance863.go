package main

/*
- Leetcode - https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/description/
- Time - O(n)
- Space - O(1)
*/

func findNodesDownwards(node *TreeNode, k int, out []int) []int {
	if node == nil || k < 0 {
		return out
	}

	if k == 0 {
		out = append(out, node.Val)
	}

	out = findNodesDownwards(node.Left, k-1, out)
	out = findNodesDownwards(node.Right, k-1, out)
	return out
}

func distanceKRecursively(node, target *TreeNode, currentDistance, targetDistance int, out []int) (int, []int) {

	if node == nil {
		return -1, out
	}

	if node == target {
		out = findNodesDownwards(node, targetDistance, out)
		return 0, out
	}

	leftNodeDistance, out := distanceKRecursively(node.Left, target, currentDistance, targetDistance, out)

	if leftNodeDistance != -1 {
		currentNodeDistance := leftNodeDistance + 1

		if currentNodeDistance == targetDistance {
			out = append(out, node.Val)
		}

		out = findNodesDownwards(node.Right, targetDistance-currentNodeDistance-1, out)
		return currentNodeDistance, out
	}

	rightNodeDistance, out := distanceKRecursively(node.Right, target, currentDistance, targetDistance, out)

	if rightNodeDistance != -1 {
		currentNodeDistance := rightNodeDistance + 1

		if currentNodeDistance == targetDistance {
			out = append(out, node.Val)
		}

		out = findNodesDownwards(node.Left, targetDistance-currentNodeDistance-1, out)
		return currentNodeDistance, out
	}

	return -1, out
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {

	out := []int{}
	_, out = distanceKRecursively(root, target, k, k, out)

	return out
}

// func main() {

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 3
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 5
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 6
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 2
// 	root.Left.Right.Left = new(TreeNode)
// 	root.Left.Right.Left.Val = 7
// 	root.Left.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Val = 4

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 1
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 0
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 8

// 	log.Println(distanceK(root, root.Left, 2))

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 20
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 8
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 4
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 12
// 	// root.Left.Right.Left = new(TreeNode)
// 	// root.Left.Right.Left.Val = 10
// 	// root.Left.Right.Right = new(TreeNode)
// 	// root.Left.Right.Right.Val = 14

// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 22

// 	// log.Println(distanceK(root, root.Left, 2))

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 0
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 1
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 3
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 2

// 	// log.Println(distanceK(root, root.Left.Right, 1))
// }
