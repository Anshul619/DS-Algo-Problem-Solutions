package main

/*
- LeetCode - https://leetcode.com/problems/basic-calculator/description/
- Time - O(n)
- Space - O(n)
*/

type calStack []int

func (s *calStack) push(r int) {
	*s = append(*s, r)
}

func (s *calStack) pop() int {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func calculate1(s string) int {
	stk := new(calStack) // stack to hold previous results and signs
	sign := 1            // current sign to next number (1 for +, -1 for -)
	ans := 0

	for i := 0; i < len(s); i++ {

		switch s[i] {
		case ' ':
			// skip spaces

		case '+':
			sign = 1

		case '-':
			sign = -1

		case '(':

			// push current result and sign to stack
			stk.push(ans)
			stk.push(sign)

			// reset for new sub-expression
			ans = 0
			sign = 1

		case ')':

			// Apply the sign first - multiply since sign (either -1 or 1) would be first to pop
			ans *= stk.pop()

			// Add the results before (
			ans += stk.pop()

		default:

			// digit detected
			num := 0
			j := i

			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				num = num*10 + int(s[j]-'0')
				j++
			}

			ans += sign * num
			i = j - 1
		}
	}

	return ans
}
