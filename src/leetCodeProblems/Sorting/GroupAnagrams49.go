package main

/*
- LeetCode - https://leetcode.com/problems/group-anagrams/description/
- Time - O(n * mlogm) (where n is length of input array and m is length of the word)
- Space - O(n * m)
*/

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, v := range strs {
		bytesArr := []byte(v)
		sort.Slice(bytesArr, func(i, j int) bool { return bytesArr[i] < bytesArr[j] })
		m[string(bytesArr)] = append(m[string(bytesArr)], v)
	}

	out := [][]string{}
	for _, v := range m {
		out = append(out, v)
	}
	return out
}
