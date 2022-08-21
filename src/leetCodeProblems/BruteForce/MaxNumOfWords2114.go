package main

/*
LeetCode - https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/submissions/
*/

import (
	"log"
	"strings"
)

func mostWordsFound(sentences []string) int {

	max := 0

	for _, v := range sentences {
		split := strings.Split(v, " ")
		if max < len(split) {
			max = len(split)
		}
	}

	return max
}

func main() {

	input := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}

	log.Println(mostWordsFound(input))
}
