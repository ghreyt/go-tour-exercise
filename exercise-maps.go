package main

import (
	//"code.google.com/p/go-tour/wc" // local package
	"golang.org/x/tour/wc" // go tour site
	"strings"
)

func WordCount(s string) map[string]int {
	Result := make(map[string]int)

	for _, f := range strings.Fields(s) {
		if _, ok := Result[f]; !ok {
			Result[f] = 0
		}

		Result[f]++
	}

	return Result
}

func main() {
	wc.Test(WordCount)
}
