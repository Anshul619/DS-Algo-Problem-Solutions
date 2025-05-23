package main

/*
- Leetcode - https://leetcode.com/problems/minimum-swaps-to-make-strings-equal/
- TimeComplexity - O(n)
- Space Complexity - O(1)
*/
func minimumSwap(s1 string, s2 string) int {
	xy, yx := 0, 0

	for i, v := range s1 {
		if v == 'x' && s2[i] == 'y' {
			xy++
		}

		if v == 'y' && s2[i] == 'x' {
			yx++
		}

	}

	if (xy+yx)%2 != 0 {
		return -1
	}

	return xy/2 + yx/2 + (xy%2)*2
}
