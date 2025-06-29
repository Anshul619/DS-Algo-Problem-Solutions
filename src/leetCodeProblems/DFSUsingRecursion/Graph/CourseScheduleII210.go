package main

/*
- LeetCode - https://leetcode.com/problems/course-schedule-ii/
- Time - O(V + E)
- Space - O(V + E)
*/

func buildGraph1(prerequisites [][]int) map[int][]int {
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
func isCycleExists1(fullyExplored []bool, visited []bool, m map[int][]int, course int, out *[]int) bool {
	// this course is already fully explored, so no cycle detected, return false
	if fullyExplored[course] {
		return false
	}

	// mark this course as visited in the current path.
	visited[course] = true

	for _, v := range m[course] {

		// if this course is already visited in the current path, then cycle exists
		// or if cycle exists in the subgraph, then return true
		if visited[v] || isCycleExists1(fullyExplored, visited, m, v, out) {
			return true
		}
	}

	visited[course] = false
	fullyExplored[course] = true

	// Append after visiting all dependencies
	*out = append(*out, course)

	return false
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	m := buildGraph1(prerequisites)

	out := []int{}

	// fullyExplored keeps track of courses that are already fully explored/globally visited
	// this is not to detect cycles, but like DP memorization to avoid re-exploring the same course again
	// Without this, the code would still work but with more time complexity
	fullyExplored := make([]bool, numCourses)

	for i := range numCourses {
		// if this course is already fully explored, no need to check again
		if !fullyExplored[i] && isCycleExists1(fullyExplored, make([]bool, numCourses), m, i, &out) {
			return []int{}
		}
	}

	// reverse out to get correct topological order
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
