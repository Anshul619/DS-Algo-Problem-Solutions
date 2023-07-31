package main

/*
- LeetCode - https://leetcode.com/problems/longest-common-subsequence/
- Time - O(m*n)
- Space - O(m*n)
*/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lcsUtil(dp [][]int, text1, text2 string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if dp[m][n] != -1 {
		return dp[m][n]
	}

	if text1[m-1] == text2[n-1] {
		dp[m][n] = 1 + lcsUtil(dp, text1, text2, m-1, n-1)
		return dp[m][n]
	}

	dp[m][n] = max(lcsUtil(dp, text1, text2, m, n-1), lcsUtil(dp, text1, text2, m-1, n))
	return dp[m][n]
}

func longestCommonSubsequence(text1 string, text2 string) int {

	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i, v := range dp {
		for j := range v {
			dp[i][j] = -1
		}
	}
	return lcsUtil(dp, text1, text2, len(text1), len(text2))
}

// func main() {
// 	log.Println(longestCommonSubsequence("abcde", "ace"))
// 	// log.Println(longestCommonSubsequence("abc", "abc"))
// 	// log.Println(longestCommonSubsequence("abc", "def"))
// 	// log.Println(longestCommonSubsequence("aggtab", "gxtxayb"))
// 	// log.Println(longestCommonSubsequence("bd", "abcd"))
// 	//log.Println(longestCommonSubsequence("mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"))
// }
