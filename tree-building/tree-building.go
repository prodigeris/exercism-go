package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	var err error

	// validation
	hasRoot := false
	ids := make(map[int]bool)

	for _, r := range records {
		if r.ID == 0 {
			hasRoot = true
		}
		if r.ID == 0 && r.Parent != 0 {
			return nil, errors.New("root node cannot have parent")
		}

		if _, exists := ids[r.ID]; exists {
			return nil, errors.New("node id cannot already exits")
		}

		if r.ID != 0 && r.Parent == r.ID {
			return nil, errors.New("cycle directly")
		}

		if r.ID < r.Parent {
			return nil, errors.New("cycle indirectly")
		}

		ids[r.ID] = true
	}

	if !hasRoot {
		return nil, errors.New("root node cannot have parent")
	}
	for i := 0; i < len(ids); i++ {
		if _, exists := ids[i]; !exists {
			return nil, errors.New("non-continuous ids")
		}
	}

	recordMap := make(map[int][]int)
	for _, r := range records {
		if v, ok := recordMap[r.Parent]; !ok {
			recordMap[r.Parent] = make([]int, len(v))
		}

		if r.Parent == r.ID {
			continue
		}

		recordMap[r.Parent] = append(recordMap[r.Parent], r.ID)
		sort.Ints(recordMap[r.Parent])
	}

	return buildNode(0, recordMap), err
}

func buildNode(key int, records map[int][]int) *Node {
	if c, ok := records[key]; ok {
		children := make([]*Node, 0)
		for _, v := range c {
			children = append(children, buildNode(v, records))
		}
		return &Node{ID: key, Children: children}
	} else {
		return &Node{ID: key}
	}
}
