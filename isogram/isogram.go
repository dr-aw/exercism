package isogram

import "strings"

func IsIsogram(word string) bool {
	var m map[rune]bool
	m = make(map[rune]bool)
	word = strings.ToLower(word)
	for _, v := range word {
		if v == ' ' || v == '-' {
			continue
		}
		_, exists := m[v]
		if exists {
			return false
		} else {
			m[v] = true
			continue
		}
	}
	return true
}
