package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	inputMap := make(map[rune]bool)

	for _, v := range input {
		inputMap[v] = true
	}

	for _, v := range alphabet {
		if _, found := inputMap[v]; !found {
			return false
		}
	}
	return true
}
