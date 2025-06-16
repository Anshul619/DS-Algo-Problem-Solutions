package main

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5, // hit -> hot -> dot -> dog -> cog
		},
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0, // cog is not in wordList
		},
		{
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			expected:  2, // a -> c (direct transform)
		},
		{
			beginWord: "lost",
			endWord:   "cost",
			wordList:  []string{"most", "fost", "lost", "cost", "host"},
			expected:  2, // lost -> cost
		},
		{
			beginWord: "same",
			endWord:   "same",
			wordList:  []string{"same"},
			expected:  1, // begin == end
		},
		{
			beginWord: "hit",
			endWord:   "hit",
			wordList:  []string{"hot", "dot", "dog"},
			expected:  1, // begin == end, even if end not in wordList
		},
		{
			beginWord: "red",
			endWord:   "tax",
			wordList:  []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
			expected:  4, // red -> ted -> tad -> tax
		},
	}

	for i, tt := range tests {
		if ladderLength(tt.beginWord, tt.endWord, tt.wordList) != tt.expected {
			t.Errorf("failed %v", i)
		}
	}
}
