package main

/*
- Leetcode - https://leetcode.com/problems/word-break/description/
- Time - O(N^2)
- Space - O(N+d) where d is dictionary size
*/
func wordBreak(s string, wordDict []string) bool {

	// map of dictionary words
	m := make(map[string]bool, len(wordDict))

	for _, v := range wordDict {
		m[v] = true
	}

	// DP array where alreadyChecked[i] means s[:i] can be segmented into dictionary words
	alreadyChecked := make([]bool, len(s)+1)
	alreadyChecked[0] = true

	for i := 1; i < len(s)+1; i++ {

		// for every i, check for all substrings
		for j := 0; j < i; j++ {

			// alreadyChecked[j] true means s[:j] already checked and present in dictionary
			// m[s[j:i]] means s[j:i] is present in dictionary
			// hence alreadyChecked[i]=true i.e. s[:i] is found in dictionary
			if alreadyChecked[j] && m[s[j:i]] {
				alreadyChecked[i] = true
				break
			}
		}
	}

	return alreadyChecked[len(s)]
}
