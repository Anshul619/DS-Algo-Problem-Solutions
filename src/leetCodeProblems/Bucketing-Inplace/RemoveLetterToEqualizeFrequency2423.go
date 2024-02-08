package main

/*
- Leetcode - https://leetcode.com/problems/remove-letter-to-equalize-frequency/description/
- Time - O(n)
- Space - O(26)
*/

func checkAllFrequencySame(freq []int) bool {
	unique := 0

	for j := 0; j < 26; j++ {
		if freq[j] != 0 {
			switch {
			case unique == 0:
				unique = freq[j]
			case unique != freq[j]:
				return false
			}
		}
	}
	return true
}
func equalFrequency(word string) bool {
	freq := make([]int, 26)

	for _, v := range word {
		freq[int(v)-int('a')]++
	}

	for i := 0; i < 26; i++ {
		if freq[i] != 0 {
			freq[i]--
			if checkAllFrequencySame(freq) {
				return true
			}
			freq[i]++
		}
	}
	return false
}
