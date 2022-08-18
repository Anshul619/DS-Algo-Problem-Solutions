package main

/*
	LeetCode - https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
*/
import "log"

func subtractProductAndSum(n int) int {

	product := 1
	sum := 0

	for n != 0 {
		digit := n % 10
		n = n / 10
		product *= digit
		sum += digit
	}

	return product - sum
}

func main() {
	log.Println(subtractProductAndSum(234))
}
