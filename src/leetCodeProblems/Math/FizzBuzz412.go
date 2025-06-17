package main

import (
	"strconv"
)

/*
- LeetCode - https://leetcode.com/problems/fizz-buzz/
- Time - O(n)
- Space - O(1)
*/

func fizzBuzz(n int) []string {
	out := []string{}

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			out = append(out, "FizzBuzz")
		case i%3 == 0:
			out = append(out, "Fizz")
		case i%5 == 0:
			out = append(out, "Buzz")
		default:
			out = append(out, strconv.Itoa(i))
		}
	}
	return out
}
