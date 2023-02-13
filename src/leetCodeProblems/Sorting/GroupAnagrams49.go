package main

/*
- LeetCode - https://leetcode.com/problems/group-anagrams/description/
- Time - O(n*mlogm) (where n is length of input array and m is length of the word)
- Space - O(n)
*/

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)
	out := [][]string{}

	for _, v := range strs {

		bytesArray := []byte(v)

		sort.Slice(bytesArray, func(i, j int) bool { return bytesArray[i] < bytesArray[j] })

		if _, ok := m[string(bytesArray)]; ok {
			m[string(bytesArray)] = append(m[string(bytesArray)], v)
		} else {
			m[string(bytesArray)] = []string{v}
		}
	}

	for _, v := range m {
		out = append(out, v)
	}

	return out
}

// func main() {
// 	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

// 	log.Println(groupAnagrams(strs))
// }
