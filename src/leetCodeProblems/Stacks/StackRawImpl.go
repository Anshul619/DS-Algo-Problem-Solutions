package main

//import "log"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}
func (s *Stack) Len() int {
	return len(*s)
}
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		elem := (*s)[index]
		*s = (*s)[:index]
		return elem, true
	}
}

// func main() {

// 	stack := new(Stack)
// 	stack.Push("One")
// 	stack.Push("Two")
// 	log.Println(stack.Pop())
// 	log.Println(stack.Pop())
// }
