package main

/*
- LeetCode - https://leetcode.com/problems/goal-parser-interpretation/description/
- Time - O(n)
- Space - O(1)
*/
import "log"

func interpret(command string) string {
	out := []rune{}
	isPrevOpenBracket := false

	for _, v := range command {
		switch string(v) {
		case "G":
			out = append(out, v)
		case "(":
			isPrevOpenBracket = true
		case ")":
			if isPrevOpenBracket {
				out = append(out, rune('o'))
			}
			isPrevOpenBracket = false
		case "a":
			if isPrevOpenBracket {
				out = append(out, rune('a'))
				out = append(out, rune('l'))
			}
			isPrevOpenBracket = false
		}
	}

	return string(out)
}

func main() {
	log.Println(interpret("G()(al)"))
	log.Println(interpret("G()()()()(al)"))
	log.Println(interpret("(al)G(al)()()G"))
}
