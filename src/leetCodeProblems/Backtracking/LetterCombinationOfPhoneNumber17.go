package main

/*
- LeetCode - https://leetcode.com/problems/letter-combinations-of-a-phone-number/
*/

func dfs1(m map[string][]string, digits string, out *[]string, combination string, index int) {

	if index == len(digits) {
		*out = append(*out, combination)
		return
	}

	for _, l := range m[string(digits[index])] {
		dfs1(m, digits, out, combination+l, index+1)
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

	dfs1(m, digits, &out, "", 0)

	return out
}

// func main() {
// 	//digits := "1234556789"
// 	digits := "2314"
// 	//digits := ""
// 	log.Println(letterCombinations(digits))
// }
