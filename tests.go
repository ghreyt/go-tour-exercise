// Codes just to check whether something unclear
package main

import (
	"fmt"
)

// can not use nil for channel for primitive types
func nilIntoPrimitiveChannel() {
	ch := make(chan string)
	ch <- nil
	fmt.Println(<-ch)
}
