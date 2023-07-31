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
	stack  *stack3
	mStack *stack3
}

func Constructor1() MinStack {
	return MinStack{
		stack:  new(stack3),
		mStack: new(stack3),
	}
}

func (this *MinStack) Push(val int) {

	if this.mStack.isEmpty() {
		this.mStack.push(val)
	} else if this.mStack.peek() >= val {
		this.mStack.push(val)
	}

	this.stack.push(val)
}

func (this *MinStack) Pop() {

	if this.stack.pop() == this.mStack.peek() {
		this.mStack.pop()
	}
}

func (this *MinStack) Top() int {
	return this.stack.peek()
}

func (this *MinStack) GetMin() int {
	return this.mStack.peek()
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
