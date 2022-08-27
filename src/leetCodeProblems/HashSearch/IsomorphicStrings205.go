package main

/*
- LeetCode - https://leetcode.com/problems/isomorphic-strings/
*/
//import "log"

func isIsomorphic(s string, t string) bool {

	dic := make(map[byte]byte)
	dic1 := make(map[byte]byte)

	for i, _ := range t {

		if mapVal, ok := dic[t[i]]; ok && mapVal != s[i] {
			return false
		} else {
			dic[t[i]] = s[i]
		}

		if mapVal, ok := dic1[s[i]]; ok && mapVal != t[i] {
			return false
		} else {
			dic1[s[i]] = t[i]
		}
	}

	return true
}

// func main() {

// 	s := "egg"
// 	t := "add"

// 	//s := "paper"
// 	//t := "title"

// 	//s := "foo"
// 	//t := "bar"

// 	//s := "badc"
// 	//t := "baba"

// 	log.Println(isIsomorphic(s, t))
// }
