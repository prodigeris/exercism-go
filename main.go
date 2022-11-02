package main

import (
	"exercism/tree-building"
	"fmt"
)

func main() {
	fmt.Println(tree.Build([]tree.Record{
		{ID: 2, Parent: 0},
		{ID: 4, Parent: 2},
		{ID: 1, Parent: 0},
		{ID: 0},
	}))
}
