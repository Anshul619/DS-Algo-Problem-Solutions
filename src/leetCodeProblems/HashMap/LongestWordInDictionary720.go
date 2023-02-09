package main

/*
- LeetCode - https://leetcode.com/problems/longest-word-in-dictionary/
*/
import (
	"log"
)

func longestWord(words []string) string {

	out := ""
	m := make(map[string]bool)

	for _, v := range words {

		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}

	for i := len(words) - 1; i >= 0; i-- {

		if (len(out) < len(words[i])) || (len(out) == len(words[i]) && words[i] < out) {

			var wordRune []rune
			found := true

			for _, v := range words[i] {
				wordRune = append(wordRune, v)

				if _, ok := m[string(wordRune)]; !ok {
					found = false
					break
				}
			}

			if found {
				out = words[i]
			}
		}
	}

	return out
}

func main() {

	//words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	//words := []string{"w", "wo", "wor", "worl", "world"}
	words := []string{"b", "br", "bre", "brea", "break", "breakf", "breakfa", "breakfas", "breakfast", "l", "lu", "lun", "lunc", "lunch", "d", "di", "din", "dinn", "dinne", "dinner"}
	log.Println("out =>", longestWord(words))
}
