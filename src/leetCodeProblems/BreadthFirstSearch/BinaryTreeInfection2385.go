package main

/*
- LeetCode - https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/description/
- Time - O(n)
- Space - O(n)
*/

type nodesQueue []*TreeNode

func (q *nodesQueue) pop() *TreeNode {
	temp := (*q)[0]
	*q = (*q)[1:]
	return temp
}

func (q *nodesQueue) push(node *TreeNode) {
	*q = append(*q, node)
}

func (q nodesQueue) isEmpty() bool {
	return len(q) == 0
}

func (q nodesQueue) Len() int {
	return len(q)
}

func findStartNode(node *TreeNode, start int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val == start {
		return node
	}

	if startNode := findStartNode(node.Left, start); startNode != nil {
		return startNode
	}

	if startNode := findStartNode(node.Right, start); startNode != nil {
		return startNode
	}

	return nil
}

func findParents(node *TreeNode, m map[int]*TreeNode) {

	if node == nil {
		return
	}

	if node.Left != nil {
		m[node.Left.Val] = node
		findParents(node.Left, m)
	}

	if node.Right != nil {
		m[node.Right.Val] = node
		findParents(node.Right, m)
	}
}

func amountOfTime(root *TreeNode, start int) int {

	parentsMap := make(map[int]*TreeNode)
	injectedMap := make(map[int]bool)
	out := -1

	startNode := findStartNode(root, start)
	findParents(root, parentsMap)

	q := new(nodesQueue)

	if startNode != nil {
		q.push(startNode)
		injectedMap[startNode.Val] = true
	}

	for !q.isEmpty() {
		size := q.Len()
		out++
		for i := 0; i < size; i++ {
			node := q.pop()
			if node.Left != nil {
				if _, ok := injectedMap[node.Left.Val]; !ok {
					q.push(node.Left)
					injectedMap[node.Left.Val] = true
				}
			}

			if node.Right != nil {
				if _, ok := injectedMap[node.Right.Val]; !ok {
					q.push(node.Right)
					injectedMap[node.Right.Val] = true
				}
			}

			if parent, ok := parentsMap[node.Val]; ok {
				if _, ok := injectedMap[parent.Val]; !ok {
					q.push(parent)
					injectedMap[parent.Val] = true
				}
			}
		}
	}

	return out
}

// func main() {
// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 1
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 5
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 4
// 	// root.Left.Right.Left = new(TreeNode)
// 	// root.Left.Right.Left.Val = 9
// 	// root.Left.Right.Right = new(TreeNode)
// 	// root.Left.Right.Right.Val = 2

// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 3
// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Left.Val = 10
// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 6

// 	// log.Println(amountOfTime(root, 3))

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 3
// 	root.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Val = 4
// 	root.Left.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Left.Val = 5

// 	log.Println(amountOfTime(root, 3))
// }
