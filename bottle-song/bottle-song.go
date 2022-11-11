package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	r := make([]string, 0)
	for i := startBottles; i > startBottles-takeDown; i-- {
		if len(r) > 0 {
			r = append(r, "")
		}
		r = append(r, verse(i)...)
	}
	return r
}

func numberToWord(n int) string {
	numbers := map[int]string{10: "ten", 9: "nine", 8: "eight", 7: "seven", 6: "six", 5: "five", 4: "four", 3: "three", 2: "two", 1: "one"}
	r, ok := numbers[n]
	if !ok {
		panic(fmt.Sprintf("%d number not found", n))
	}
	return r
}

func verse(n int) []string {
	return []string{
		firstLine(n),
		firstLine(n),
		"And if one green bottle should accidentally fall,",
		lastLine(n - 1),
	}
}

func firstLine(n int) string {
	switch n {
	case 1:
		return "One green bottle hanging on the wall,"
	default:
		return fmt.Sprintf("%s green bottles hanging on the wall,", strings.Title(numberToWord(n)))
	}
}

func lastLine(n int) string {
	switch n {
	case 1:
		return fmt.Sprintf("There'll be one green bottle hanging on the wall.")
	case 0:
		return fmt.Sprintf("There'll be no green bottles hanging on the wall.")
	default:
		return fmt.Sprintf("There'll be %s green bottles hanging on the wall.", numberToWord(n))
	}
}
