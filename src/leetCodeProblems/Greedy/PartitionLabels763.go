package main

/*
- LeetCode - https://leetcode.com/problems/partition-labels/description/
- Time - O(n)
- Space - O(1)
*/

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func partitionLabels(s string) []int {

	last := make([]int, 26)
	out := []int{}

	for i, v := range s {
		last[v-rune('a')] = i
	}

	partitionStart, partitionEnd := 0, 0

	for i, v := range s {
		partitionEnd = getMax(partitionEnd, last[v-rune('a')])

		if i == partitionEnd {
			out = append(out, partitionEnd-partitionStart+1)
			partitionStart = partitionEnd + 1
		}
	}

	return out
}

// func main() {

// 	//s := "ababcbacadefegdehijhklij"
// 	s := "eccbbbbdec"

// 	log.Println(partitionLabels(s))
// }
