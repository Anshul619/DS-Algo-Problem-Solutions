package main

/*
- LeetCode - https://leetcode.com/problems/accounts-merge/
*/
import (
	"log"
	"sort"
)

func dfs(email string, neighbors []string, emailToNeighbors map[string][]string, isVisited map[string]bool, temp *[]string) {

	// log.Println(isVisited)

	if isVisited[email] {
		return
	}

	isVisited[email] = true

	for _, v := range neighbors {

		if !isVisited[v] {
			*temp = append(*temp, v)
			dfs(v, emailToNeighbors[v], emailToNeighbors, isVisited, temp)
		}
	}
}

func accountsMerge(accounts [][]string) [][]string {

	var out [][]string

	emailToName := make(map[string]string)
	emailToNeighbors := make(map[string][]string)
	isVisited := make(map[string]bool)

	log.Println("input", accounts)

	for _, v := range accounts {

		name := v[0] // + strconv.Itoa(i)

		for j := 1; j < len(v); j++ {
			emailToName[v[j]] = name

			if _, ok := emailToNeighbors[v[j]]; !ok {
				emailToNeighbors[v[j]] = []string{}
			}

			for k := 1; k < len(v); k++ {
				if j != k {
					emailToNeighbors[v[j]] = append(emailToNeighbors[v[j]], v[k])
				}
			}
		}
	}

	log.Println("emailToName", emailToName)
	log.Println("emailToNeighbors", emailToNeighbors)

	for email, neighbors := range emailToNeighbors {

		// log.Println(email, isVisited[email])

		if !isVisited[email] {
			temp := []string{email}

			dfs(email, neighbors, emailToNeighbors, isVisited, &temp)

			if len(temp) > 0 {
				sort.Strings(temp)
				out = append(out, append([]string{emailToName[email]}, temp...))
			}
		}
	}

	return out
}

// func main() {
// 	input := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
// 	//input := [][]string{{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"}, {"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"}, {"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"}, {"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"}, {"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"}}
// 	//input := [][]string{{"David", "David0@m.co", "David4@m.co", "David3@m.co"}, {"David", "David5@m.co", "David5@m.co", "David0@m.co"}, {"David", "David1@m.co", "David4@m.co", "David0@m.co"}, {"David", "David0@m.co", "David1@m.co", "David3@m.co"}, {"David", "David4@m.co", "David1@m.co", "David3@m.co"}}

// 	//input := [][]string{{"David", "David0@m.co", "David1@m.co"}, {"David", "David3@m.co", "David4@m.co"}, {"David", "David4@m.co", "David5@m.co"}, {"David", "David2@m.co", "David3@m.co"}, {"David", "David1@m.co", "David2@m.co"}}
// 	log.Println("output", accountsMerge(input))
// }
