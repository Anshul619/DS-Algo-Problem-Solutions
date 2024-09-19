package main

/*
- Leetcode - https://leetcode.com/problems/basic-calculator/description/?envType=study-plan-v2&envId=top-interview-150
*/

// type stack []string

// func (s *stack) push(r string) {
// 	*s = append(*s, r)
// }

// func (s *stack) pop() string {
// 	t := (*s)[len(*s)-1]
// 	*s = (*s)[:len(*s)-1]
// 	return t
// }

// func (s stack) peek() string {
// 	return s[len(s)-1]
// }

// func (s stack) isEmpty() bool {
// 	return len(s) == 0
// }

// func (s stack) len() int {
// 	return len(s)
// }

// func eval(op, n1 string, n2 int) (int, error) {
// 	num1, err := strconv.Atoi(n1)
// 	if err != nil {
// 		i, _ := strconv.Atoi(op + strconv.Itoa(n2))
// 		return i, err
// 	}

// 	switch op {
// 	case "+":
// 		return num1 + n2, nil
// 	case "-":
// 		return num1 - n2, nil
// 	default:
// 		i, _ := strconv.Atoi(n1 + op + strconv.Itoa(n2))
// 		return i, nil
// 	}
// }

// func calculate(s string) int {
// 	st := new(stack)
// 	open := 0

// 	for _, v := range s {
// 		switch string(v) {
// 		case "(":
// 			open++
// 			st.push(string(v))
// 		case "+", "-":
// 			st.push(string(v))
// 		case " ":
// 			continue
// 		case ")":
// 			log.Println(")", st, string(v))
// 			res, _ := strconv.Atoi(st.pop())

// 			for !st.isEmpty() && st.peek() != "(" {
// 				res1, err := eval(st.pop(), st.pop(), res)
// 				res = res1
// 				log.Println("err", err)
// 				if err != nil {

// 					break
// 				}
// 			}

// 			if st.peek() == "(" {
// 				st.pop()
// 			}

// 			st.push(strconv.Itoa(res))
// 			open--
// 			log.Println(") after", st, string(v))
// 		default:
// 			log.Println("num", st, string(v))
// 			if !st.isEmpty() && open == 0 {
// 				if st.len() > 1 {
// 					res, _ := strconv.Atoi(string(v))
// 					res, _ = eval(st.pop(), st.pop(), res)
// 					st.push(strconv.Itoa(res))
// 				} else {
// 					st.push(st.pop() + string(v))
// 				}

// 			} else {
// 				st.push(string(v))
// 			}
// 			log.Println("num after", st, string(v))
// 		}
// 	}

// 	log.Println(st)
// 	if len(*st) == 3 {
// 		log.Println("st", st)
// 		res, _ := strconv.Atoi(st.pop())
// 		res, _ = eval(st.pop(), st.pop(), res)
// 		return res
// 	}

// 	if len(*st) == 2 {
// 		n1 := st.pop()
// 		i, _ := strconv.Atoi(st.pop() + n1)
// 		return i
// 	}
// 	out, _ := strconv.Atoi(st.pop())
// 	return out
// }
