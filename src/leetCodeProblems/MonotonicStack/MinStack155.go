package main

/*
- Leetcode - https://leetcode.com/problems/min-stack/description/
- Time - O(1)
- Space - O(n)
*/

type stack3 []int

func (s *stack3) push(n int) {
	*s = append(*s, n)
}

func (s *stack3) pop() int {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s stack3) isEmpty() bool {
	return len(s) == 0
}

func (s stack3) peek() int {
	return s[len(s)-1]
}

type MinStack struct {
	s  *stack3
	ms *stack3
}

func Constructor1() MinStack {
	return MinStack{
		s:  new(stack3),
		ms: new(stack3),
	}
}

func (this *MinStack) Push(val int) {
	if this.ms.isEmpty() || this.ms.peek() >= val {
		this.ms.push(val)
	}

	this.s.push(val)
}

func (this *MinStack) Pop() {
	if !this.ms.isEmpty() &&
		this.s.peek() == this.ms.peek() {
		this.ms.pop()
	}

	this.s.pop()
}

func (this *MinStack) Top() int {
	return this.s.peek()
}

func (this *MinStack) GetMin() int {
	return this.ms.peek()
}

// func main() {
// 	min := Constructor1()
// 	min.Push(-2)
// 	min.Push(0)
// 	min.Push(-3)
// 	log.Println(min.GetMin())
// 	min.Pop()
// 	log.Println(min.Top())
// 	log.Println(min.GetMin())
// }
