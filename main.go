package main

import (
	account "exercism/bank-account"
	"fmt"
)

func main() {
	a := account.Open(100)
	a.Close()
	fmt.Println(a.Balance())
}
