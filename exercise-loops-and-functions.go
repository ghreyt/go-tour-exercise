package main

import (
	"fmt"
	"math"
)

const (
	MAX = 10      // max calculation
	ACC = 1000000 // accuracy
)

// Newton's method for Sqrt which limits accuracy with ACC
func NTMethod(x, z float64) float64 {
	return math.Ceil((z-(z*z-x)/(2*z))*ACC) / ACC
}

// try to find initial value as near to the answer as possible
func PickInitVal(x float64) float64 {
	// a. simply returns just half of x
	//return x / 2

	// b. try to find half bit-length number
	// (all set 0 except the first, 1000..00)
	v := float64(1)
	for v*v < x {
		v *= 2
	}
	return v
}

func Sqrt(x float64) float64 {
	// pick initial value
	z := PickInitVal(x)

	// caculate by Newton's method
	i := 0
	prev := z
	for ; i < MAX; i++ {
		if z = NTMethod(x, z); z == prev {
			// stop if z is the same as before
			break
		}
		prev = z
	}
	fmt.Print("sqrt(", x, ") found in ", i, " times = ")

	return z
}

func main() {
	fmt.Println(Sqrt(3319))
}
