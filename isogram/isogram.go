package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for i, v := range word {
		if v == ' ' || v == '-' {
			continue
		}
		if strings.ContainsRune(word[i+1:], v) {
			return false
		}
	}
	return true
}
