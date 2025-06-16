package main

/*
- Leetcode - https://leetcode.com/problems/word-ladder/description
- Time - O(N * L^2) where N is number of words, L is word length
- Space - O(N*L) - queue + map
*/

type queue []string

func (q *queue) push(i string) {
	*q = append(*q, i)
}

func (q *queue) pop() string {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}

func (q queue) isEmpty() bool {
	return len(q) == 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := make(map[string]bool, len(wordList))

	for _, v := range wordList {
		m[v] = true
	}

	if beginWord == endWord {
		return 1
	}

	if _, ok := m[endWord]; !ok {
		return 0
	}

	out := 1

	q := new(queue)
	q.push(beginWord)

	for !q.isEmpty() {
		q1 := new(queue)

		for !q.isEmpty() {

			s := q.pop()

			if s == endWord {
				return out
			}

			for i := 0; i < len(s); i++ {
				for c := 'a'; c <= 'z'; c++ {
					if rune(s[i]) == c {
						continue
					}

					newWord := s[0:i] + string(c) + s[i+1:]

					if _, ok := m[newWord]; ok {
						q1.push(newWord)
						delete(m, newWord)
					}
				}
			}
		}

		q = q1
		out++
	}

	return 0
}
