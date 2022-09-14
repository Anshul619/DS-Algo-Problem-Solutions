package main

/*
- LeetCode - https://leetcode.com/problems/letter-combinations-of-a-phone-number/
*/
import (
	"log"
)

func util(digits string, index int, m map[string][]string, out *[]string, combination string) {

	if index == len(digits) {
		*out = append(*out, combination)
		return
	}

	sv := string(digits[index])

	if _, ok := m[sv]; ok {
		for _, letter := range m[sv] {
			util(digits, index+1, m, out, combination+letter)
		}
	} else {
		util(digits, index+1, m, out, combination)
	}
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	m := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	out := []string{}

	util(digits, 0, m, &out, "")

	return out
}

func main() {
	//digits := "1234556789"
	digits := "2314"
	//digits := ""
	log.Println(letterCombinations(digits))
}
