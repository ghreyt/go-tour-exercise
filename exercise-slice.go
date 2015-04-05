package main

import "golang.org/x/tour/pic"

func Trans(x, y int) uint8 {
	// any of below
	//return uint8((x + y) / 2)
	//return uint8(x ^ y)
	return uint8(x * y)
}

func Pic(dx, dy int) [][]uint8 {
	Result := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		Result[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			Result[y][x] = Trans(x, y)
		}
	}

	return Result
}

func main() {
	pic.Show(Pic)
}
