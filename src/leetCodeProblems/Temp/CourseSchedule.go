package main

/*
- Leetcode - https://leetcode.com/problems/course-schedule/description/?envType=study-plan-v2&envId=top-interview-150
*/
import "log"

type queue []int

func (q *queue) pop() int {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}

func (q *queue) push(e int) {
	*q = append(*q, e)
}

func (q queue) isEmpty() bool {
	return len(q) == 0
}

func buildGraph(numCourses int, prerequisites [][]int) map[int][]int {
	m := make(map[int][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		m[i] = []int{}
	}

	for _, v := range prerequisites {
		m[v[1]] = append(m[v[1]], v[0])
	}
	return m
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	m := buildGraph(numCourses, prerequisites)
	q := new(queue)

	q.push(prerequisites[0][1]) // add first course prerequisite as starting point

	visited := make(map[int]bool, numCourses)

	log.Println(m, visited)

	for !q.isEmpty() {
		v := q.pop()

		log.Println(visited)
		// cycle detected
		if visited[v] {
			continue
		}

		for _, v1 := range m[v] {
			q.push(v1)
		}
		visited[v] = true
	}

	//log.Println(m, visited)
	// Check if all courses are completed
	for i := 0; i < numCourses; i++ {
		if !visited[i] && len(m[i]) != 0 {
			return false
		}
	}
	return true
}
