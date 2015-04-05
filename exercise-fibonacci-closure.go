package main

import "fmt"

func fibonacci() func() int {
	prev, curr := 1, 0
	return func() int {
		prev, curr = curr, prev+curr
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
