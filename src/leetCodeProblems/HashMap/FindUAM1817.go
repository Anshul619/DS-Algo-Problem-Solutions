package main

/*
- LeetCode - https://leetcode.com/problems/finding-the-users-active-minutes/description/
- Time - O(n)
- Space - O(n)
*/

func findingUsersActiveMinutes(logs [][]int, k int) []int {

	m := make(map[int]map[int]bool)
	out := make([]int, k)

	for _, v := range logs {

		if _, ok := m[v[0]]; ok {

			if _, ok1 := m[v[0]][v[1]]; !ok1 {
				m[v[0]][v[1]] = true
			}
		} else {
			m[v[0]] = make(map[int]bool)
			m[v[0]][v[1]] = true
		}
	}

	for _, v := range m {
		out[len(v)-1]++
	}

	return out
}

// func main() {
// 	logs := [][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}

// 	log.Println(findingUsersActiveMinutes(logs, 5))
// }
