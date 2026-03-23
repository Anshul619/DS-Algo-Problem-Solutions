package main

/*
- Leetcode - https://leetcode.com/problems/all-paths-from-source-to-target/
*/
func recurPath(graph [][]int, path *[]int, out *[][]int, node int) {

	*path = append(*path, node)

	// reached target
	if node == len(graph)-1 {
		*out = append(*out, *path)
		return
	}

	for _, v := range graph[node] {
		newPath := []int{}
		newPath = append(newPath, *path...)

		recurPath(graph, &newPath, out, v)
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	out := [][]int{}

	// 0 is the source
	for _, v := range graph[0] {
		path := []int{0}
		recurPath(graph, &path, &out, v)
	}
	return out
}
