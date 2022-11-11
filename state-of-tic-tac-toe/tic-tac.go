package stateoftictactoe

import (
	"errors"
	"fmt"
	"strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	var x string
	var o string
	fmt.Println(board)
	for r, v := range board {
		for c, v := range v {
			switch v {
			case 'X':
				x = fmt.Sprintf("%s%d", x, r*3+c)
			case 'O':
				o = fmt.Sprintf("%s%d", o, r*3+c)
			}
		}
	}

	if !isValid(x, o) {
		return State(""), errors.New("invalid board")
	}

	if hasWin(x) || hasWin(o) {
		return Win, nil
	}

	if len(x)+len(o) < 9 {
		return Ongoing, nil
	}

	return Draw, nil
}

func isValid(x string, o string) bool {
	return len(x) >= len(o) && len(x)-len(o) < 2 && !(hasWin(x) && hasWin(o))
}

func hasWin(ticks string) bool {
	wins := []string{"012", "345", "678", "036", "147", "258", "048", "246", "0348"}
	for _, v := range wins {
		if strings.Contains(ticks, v) {
			return true
		}
	}
	return false
}
