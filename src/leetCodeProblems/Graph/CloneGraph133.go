package main

/*
- LeetCode - https://leetcode.com/problems/clone-graph/description/
- Time - O(n)
- Space - O(n)
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func visitNode(node *Node, m map[*Node]*Node, visted map[*Node]bool) {

	if isVisited, ok := visted[node]; isVisited && ok {
		return
	}

	visted[node] = true

	for _, v := range node.Neighbors {

		if _, ok := m[v]; !ok {
			m[v] = new(Node)
			m[v].Val = v.Val
			m[v].Neighbors = []*Node{}
		}

		m[node].Neighbors = append(m[node].Neighbors, m[v])
		visitNode(v, m, visted)
	}
}

func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	m := make(map[*Node]*Node)
	visited := make(map[*Node]bool)

	m[node] = new(Node)
	m[node].Val = node.Val
	m[node].Neighbors = []*Node{}

	visitNode(node, m, visited)

	return m[node]
}

// func main() {

// 	node1 := new(Node)
// 	node1.Val = 1
// 	node1.Neighbors = []*Node{}

// 	node2 := new(Node)
// 	node2.Val = 2
// 	node2.Neighbors = []*Node{}

// 	node3 := new(Node)
// 	node3.Val = 3
// 	node3.Neighbors = []*Node{}

// 	node4 := new(Node)
// 	node4.Val = 4
// 	node4.Neighbors = []*Node{}

// 	node1.Neighbors = append(node1.Neighbors, node2, node4)
// 	node2.Neighbors = append(node2.Neighbors, node1, node3)
// 	node3.Neighbors = append(node3.Neighbors, node2, node4)
// 	node4.Neighbors = append(node4.Neighbors, node1, node3)

// 	new := cloneGraph(node1)

// 	log.Println(new)

// 	for _, v := range new.Neighbors {
// 		log.Println(v)
// 	}
// }
