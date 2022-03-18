package main

import (
	"fmt"
)

type Node[T any] struct {
	Content T
	Next    *Node[T]
}

func main() {
	x := Node[string]{
		Content: "x",
		Next:    nil,
	}

	y := Node[string]{
		Content: "y",
		Next:    nil,
	}

	z := Node[string]{
		Content: "z",
		Next:    nil,
	}

	x.Next = &y
	y.Next = &z

	current := &x

	for current != nil {
		fmt.Println(current.Content)
		current = current.Next
	}
}

