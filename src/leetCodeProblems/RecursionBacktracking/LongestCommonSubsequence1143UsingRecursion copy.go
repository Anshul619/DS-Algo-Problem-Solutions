package main

/*
- LeetCode - https://leetcode.com/problems/longest-common-subsequence/
- Time - O(2m*n)
- Space - O(1)
*/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lcsUtil(text1, text2 string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if text1[m-1] == text2[n-1] {
		return 1 + lcsUtil(text1, text2, m-1, n-1)
	} else {
		return max(lcsUtil(text1, text2, m, n-1), lcsUtil(text1, text2, m-1, n))
	}
}

func longestCommonSubsequenceRecursion(text1 string, text2 string) int {
	return lcsUtil(text1, text2, len(text1), len(text2))
}

// func main() {
// 	// log.Println(longestCommonSubsequence("abcde", "ace"))
// 	// log.Println(longestCommonSubsequence("abc", "abc"))
// 	// log.Println(longestCommonSubsequence("abc", "def"))
// 	// log.Println(longestCommonSubsequence("aggtab", "gxtxayb"))
// 	// log.Println(longestCommonSubsequence("bd", "abcd"))
// 	log.Println(longestCommonSubsequence("mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"))
// }
