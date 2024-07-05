package main

/*
- LeetCode - https://leetcode.com/problems/clone-graph/description/
- Time - O(V + E)
- Space - O(V)
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func visit(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if visited[node.Val] != nil {
		return visited[node.Val]
	}

	newNode := new(Node)
	newNode.Val = node.Val
	newNode.Neighbors = []*Node{}

	visited[node.Val] = newNode

	for _, v := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, visit(v, visited))
	}
	return newNode
}

func cloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	return visit(node, visited)
}
