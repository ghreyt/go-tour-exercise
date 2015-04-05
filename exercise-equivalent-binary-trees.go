package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

const (
	LENGTH = 10
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	result := true
	for i := 0; i < LENGTH; i++ {
		if <-ch1 != <-ch2 {
			// not return false not to block channel
			// (so left goroutine work)
			result = false
		}
	}

	return result
}

func main() {
	fmt.Println("should be true :", Same(tree.New(1), tree.New(1)))
	fmt.Println("should be false:", Same(tree.New(1), tree.New(2)))
}
