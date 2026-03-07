package main

import "strings"

/*
- Leetcode - https://leetcode.com/problems/simplify-path/description/
- Time - O(n)
- Space - O(n)
*/
func simplifyPath(path string) string {
	stack := []string{}

	for _, v := range strings.Split(path, "/") {
		switch v {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, v)
		}
	}

	return "/" + strings.Join(stack, "/")
}
