package main

/*
- LeetCode - https://leetcode.com/problems/course-schedule-ii/
*/
import "log"

func buildDependenciesMap1(prerequisites [][]int) map[int][]int {

	dep := map[int][]int{}

	for _, val := range prerequisites {

		if _, ok := dep[val[1]]; !ok {
			dep[val[1]] = []int{}
		}
		dep[val[1]] = append(dep[val[1]], val[0])
	}

	return dep
}

func isDFSCycleExists1(dep map[int][]int, visitedOnPath []bool, alreadyVisited []bool, course int) bool {

	if alreadyVisited[course] {
		return false
	}

	alreadyVisited[course] = true
	visitedOnPath[course] = true

	for _, v := range dep[course] {
		log.Println("Course -", v)
		if visitedOnPath[v] || isDFSCycleExists1(dep, visitedOnPath, alreadyVisited, v) {
			return true
		}
	}

	visitedOnPath[course] = false

	return false
}

func canFinish1(numCourses int, prerequisites [][]int) bool {
	dep := buildDependenciesMap1(prerequisites)

	alreadyVisited := make([]bool, numCourses)
	visitedOnPath := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		log.Println(i)
		log.Println(visitedOnPath)
		if !alreadyVisited[i] && isDFSCycleExists1(dep, visitedOnPath, alreadyVisited, i) {
			return false
		}

	}

	return true
}

func main() {

	numCourses := 4
	//prerequisites := [][]int{{1, 0}}

	//prerequisites := [][]int{{1, 0}, {0, 1}}
	prerequisites := [][]int{{1, 0}, {2, 1}, {3, 1}, {3, 2}}

	log.Println(canFinish1(numCourses, prerequisites))
}
