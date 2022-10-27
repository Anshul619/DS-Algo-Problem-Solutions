package main

import (
	"log"
	"strconv"
	"strings"
)

type stackNums []int

func (s *stackNums) push(v int) {
	*s = append(*s, v)
}

func (s *stackNums) pop() int {
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out
}

func (s stackNums) isEmpty() bool {
	return len(s) == 0
}

func isDigit(s string) bool {
	if s != "/" && s != "*" && s != "+" && s != "-" {
		return true
	}

	return false
}

func calculate(s string) int {

	s = strings.ReplaceAll(s, " ", "")

	stack := new(stackNums)
	out := 0
	num := 0
	sign := string('+')

	for i := 0; i < len(s); i++ {

		v := string(s[i])

		if isDigit(v) {
			// log.Println("Digit", v)
			v3, _ := strconv.Atoi(string(s[i]))
			num = num*10 + v3
			// log.Println(num)
		}

		if i+1 == len(s) || !isDigit(v) {

			if sign == "/" {
				stack.push(stack.pop() / num)
			} else if sign == "*" {
				stack.push(stack.pop() * num)
			} else if sign == "-" {
				stack.push(-num)
			} else if sign == "+" {
				stack.push(num)
			}

			sign = v
			num = 0
		}
		// log.Println("v", v)
		// log.Println("num", num)
		// log.Println(stack, "----")
	}

	//log.Println(stack)

	for !stack.isEmpty() {
		out += stack.pop()
	}

	return out
}

func main() {

	//s := "3+2*2"
	//s := " 3/2 "
	//s := " 3+5 / 2 "
	s := "42"
	log.Println(calculate(s))
}
