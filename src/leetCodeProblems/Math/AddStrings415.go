package main

import (
	"log"
	"strconv"
)

/*
- LeetCode - https://leetcode.com/problems/add-strings/
- Time - O(n)
- Space - O(1)
*/

func addStrings(num1 string, num2 string) string {

	carry := 0
	out := ""
	i1 := len(num1) - 1
	i2 := len(num2) - 1

	for i1 >= 0 || i2 >= 0 {

		sum := carry

		if i1 >= 0 {
			num, _ := strconv.Atoi(string(num1[i1]))
			sum += num
		}

		if i2 >= 0 {
			num, _ := strconv.Atoi(string(num2[i2]))
			sum += num
		}

		out = strconv.Itoa(sum%10) + out
		carry = sum / 10
		i2--
		i1--
	}

	if carry > 0 {
		out = strconv.Itoa(carry) + out
	}
	return out
}

func main() {

	log.Println(addStrings("11", "123"))
	log.Println(addStrings("456", "77"))
}
