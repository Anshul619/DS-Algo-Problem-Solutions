package main

/*
- LeetCode - https://leetcode.com/problems/course-schedule/
*/
import "log"

func buildDependenciesMap(prerequisites [][]int) map[int][]int {

	dep := map[int][]int{}

	for _, val := range prerequisites {

		if _, ok := dep[val[1]]; !ok {
			dep[val[1]] = []int{}
		}
		dep[val[1]] = append(dep[val[1]], val[0])
	}

	return dep
}

func isDFSCycleExists(dep map[int][]int, visitedOnPath []bool, alreadyVisited []bool, course int) bool {

	if alreadyVisited[course] {
		return false
	}

	alreadyVisited[course] = true
	visitedOnPath[course] = true

	for _, v := range dep[course] {
		if visitedOnPath[v] || isDFSCycleExists(dep, visitedOnPath, alreadyVisited, v) {
			return true
		}
	}

	visitedOnPath[course] = false

	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	dep := buildDependenciesMap(prerequisites)

	alreadyVisited := make([]bool, numCourses)
	visitedOnPath := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !alreadyVisited[i] && isDFSCycleExists(dep, visitedOnPath, alreadyVisited, i) {
			return false
		}
	}

	return true
}

func main() {

	numCourses := 2
	//prerequisites := [][]int{{1, 0}}

	prerequisites := [][]int{{1, 0}, {0, 1}}

	log.Println(canFinish(numCourses, prerequisites))
}
