package main

import (
	"strconv"
)

/*

PayPay OA (Code Signal) - Question 2
- Time - O(n)
- Space - O(1)

Similar Problem - https://leetcode.com/problems/add-strings/description/

You are given two numerical strings, and your task is to return the sum of their digits, as described below.
Add every ith digit of the first string to the ith digit of the second string, both counted from the end. If the ith digit of one of the strings is absent, the sum will be the ith digit of the other string. Return a string of those sums concatenated with each other.
Note: You are not expected to provide the most optimal solution, but a solution with time complexity not worse than O(max(a.length, b.length)2) will fit within the execution time limit.
Example
* For a = "99" and b = "99", the output should be solution(a, b) = "1818".The sums of both, the first and the second numbers are 18, so the answer is 1818.
* For a = "11" and b = "9", the output should be solution(a, b) = "110".The sum of the first numbers from the end is 10, and the sum of the second numbers from the end is 1, so the answer is 110.

Note - As inputs are very big integers, we can't convert inputs to integer, then iterate and sum it.
*/

func solution2(a string, b string) string {
	ai := len(a) - 1
	bi := len(b) - 1

	out := ""

	for ai >= 0 || bi >= 0 {
		n1, n2 := 0, 0

		if ai >= 0 {
			n1, _ = strconv.Atoi(string(a[ai]))
			ai--
		}
		if bi >= 0 {
			n2, _ = strconv.Atoi(string(b[bi]))
			bi--
		}

		out = strconv.Itoa(n1+n2) + out

	}
	return out
}
