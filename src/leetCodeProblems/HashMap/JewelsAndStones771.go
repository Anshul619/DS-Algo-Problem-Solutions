package main

/*
- LeetCode - https://leetcode.com/problems/jewels-and-stones/submissions/968606888/
- Time - O(m)
- Space - O(n)
*/
import "log"

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[rune]bool)
	count := 0

	for _, v := range jewels {
		m[v] = true
	}

	for _, v := range stones {
		if _, ok := m[v]; ok {
			count++
		}
	}
	return count
}

func main() {
	log.Println(numJewelsInStones("aA", "aAAbbbb"))
	log.Println(numJewelsInStones("z", "ZZ"))
}
