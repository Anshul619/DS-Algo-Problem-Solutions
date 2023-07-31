package main

/**
- LeetCode - https://leetcode.com/problems/daily-temperatures
*/

type TempWithIndex struct {
	Val   int
	Index int
}

type stack2 []TempWithIndex

func (s *stack2) push(e TempWithIndex) {
	*s = append(*s, e)
}

func (s *stack2) size() int {
	return len(*s)
}

func (s *stack2) isEmpty() bool {
	return s.size() == 0
}

func (s *stack2) peek() (bool, TempWithIndex) {
	if !s.isEmpty() {
		return true, (*s)[len(*s)-1]
	}
	return false, TempWithIndex{}
}

func (s *stack2) pop() (bool, TempWithIndex) {
	if !s.isEmpty() {
		e := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return true, e
	}
	return false, TempWithIndex{}
}

func dailyTemperatures(temperatures []int) []int {

	out := make([]int, len(temperatures))
	monoStack := new(stack2)

	for i := len(temperatures) - 1; i > -1; i-- {

		for !monoStack.isEmpty() {
			_, elem := monoStack.peek()

			if elem.Val > temperatures[i] {
				break
			}

			monoStack.pop()
		}

		if monoStack.isEmpty() {
			out[i] = 0
		} else {
			_, elem := monoStack.peek()
			out[i] = elem.Index - i
		}

		monoStack.push(TempWithIndex{temperatures[i], i})
	}

	return out
}

// func main() {

// 	//temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
// 	//temperatures := []int{30, 40, 50, 60}
// 	temperatures := []int{30, 60, 90}

// 	log.Println(dailyTemperatures(temperatures))

// }
