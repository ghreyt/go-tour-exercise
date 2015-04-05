package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	// simply just fill one 'A'
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}