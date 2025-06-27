package main

/*
- LeetCode - https://leetcode.com/problems/course-schedule/
- Time - O(V + E)
- Space - O(V + E)
*/

func buildGraph(prerequisites [][]int) map[int][]int {
	m := map[int][]int{}

	for _, v := range prerequisites {
		m[v[1]] = append(m[v[1]], v[0])
	}
	return m
}

// fullyExplored - keeps track of courses that are already fully explored/globally visited
// visited - keeps track of courses that are visited in the current path
// m - adjacency list representation of the graph
// course - the current course being explored
func isCycleExists(fullyExplored []bool, visited []bool, m map[int][]int, course int) bool {
	// this course is already fully explored, so no cycle detected, return false
	if fullyExplored[course] {
		return false
	}

	// mark this course as visited in the current path.
	visited[course] = true

	for _, v := range m[course] {

		// if this course is already visited in the current path, then cycle exists
		// or if cycle exists in the subgraph, then return true
		if visited[v] || isCycleExists(fullyExplored, visited, m, v) {
			return true
		}
	}

	visited[course] = false
	fullyExplored[course] = true
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := buildGraph(prerequisites)

	// fullyExplored keeps track of courses that are already fully explored/globally visited
	// this is not to detect cycles, but like DP memorization to avoid re-exploring the same course again
	// Without this, the code would still work but with more time complexity
	fullyExplored := make([]bool, numCourses)

	for i := range numCourses {
		// if this course is already fully explored, no need to check again
		if !fullyExplored[i] && isCycleExists(fullyExplored, make([]bool, numCourses), m, i) {
			return false
		}
	}
	return true
}
