package main

/*
- LeetCode - https://leetcode.com/problems/asteroid-collision/
*/

import "log"

type stack []int

func (s *stack) push(elem int) {
	*s = append(*s, elem)
}

func (s *stack) peek() (bool, int) {
	if s.isEmpty() {
		return false, -1
	}

	return true, (*s)[len(*s)-1]
}

func (s *stack) pop() (bool, int) {

	if s.isEmpty() {
		return false, -1
	}

	elem := (*s)[len(*s)-1]

	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	} else {
		*s = []int{}
	}

	return true, elem
}

func (s stack) isEmpty() bool {
	return len(s) <= 0
}

func (s stack) size() int {
	return len(s)
}

func getAbsValue(num int) (string, int) {

	if num > 0 {
		return "right", num
	} else {
		return "left", -num
	}
}

func asteroidCollision(asteroids []int) []int {

	remainingAsteriods := new(stack)

	for _, v := range asteroids {

		curNumDirection, curNumAbsVal := getAbsValue(v)
		insertCurNum := false

		if curNumDirection == "right" {
			insertCurNum = true
		} else {

			if remainingAsteriods.isEmpty() {
				insertCurNum = true
			}

			for !remainingAsteriods.isEmpty() {

				_, sElem := remainingAsteriods.peek()
				sDirection, sAbsVal := getAbsValue(sElem)

				if sDirection == curNumDirection {
					insertCurNum = true
					break
				}

				if curNumAbsVal <= sAbsVal {
					insertCurNum = false

					if curNumAbsVal == sAbsVal {
						remainingAsteriods.pop()
					}

					break
				} else {
					remainingAsteriods.pop()
					insertCurNum = true
				}
			}
		}

		if insertCurNum {
			remainingAsteriods.push(v)
		}
	}

	size := remainingAsteriods.size()
	out := make([]int, size)

	counter := size - 1

	for !remainingAsteriods.isEmpty() {
		_, elem := remainingAsteriods.pop()

		out[counter] = elem
		counter--
	}

	return out
}

func main() {

	input := []int{1, -2, -2, -2}

	log.Println(asteroidCollision(input))
}
