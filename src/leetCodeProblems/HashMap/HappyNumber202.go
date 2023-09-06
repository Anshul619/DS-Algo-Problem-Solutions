package main

import "log"

/*
- Leetcode - https://leetcode.com/problems/happy-number/
- Time - O(nlogn)
- Space - O(n)
*/

func calculateSquaresSum(n int) int {
	sum := 0

	for n != 0 {
		r := n % 10
		sum += r * r

		n = n / 10
	}
	return sum
}
func isHappy(n int) bool {
	m := make(map[int]struct{})

	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}

		m[n] = struct{}{}
		n = calculateSquaresSum(n)
	}
	return n == 1
}

func main() {
	log.Println(isHappy(19))
	log.Println(isHappy(2))
	//log.Println(intToRoman(1994))
}
