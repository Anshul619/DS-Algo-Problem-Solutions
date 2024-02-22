package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)

	for _, v := range magazine {
		m[v-rune('a')]++
	}

	for _, v := range ransomNote {
		if m[v-rune('a')] <= 0 {
			return false
		}
		m[v-rune('a')]--
	}
	return true
}
