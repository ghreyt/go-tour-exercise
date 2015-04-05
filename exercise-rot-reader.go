package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(ch byte) byte {
	switch {
	case 65 <= ch && ch <= 90:
		ch = (ch-65+13)%26 + 65
	case 97 <= ch && ch <= 122:
		ch = (ch-97+13)%26 + 97
	}

	return ch
}

func (this rot13Reader) Read(b []byte) (int, error) {
	if n, err := this.r.Read(b); err == nil {
		for i, ch := range b {
			b[i] = rot13(ch)
		}

		return n, nil
	} else {
		return 0, err
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
