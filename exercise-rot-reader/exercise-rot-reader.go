// What is rot13?
// Rot13 is a simple letter substitution cipher that replaces a letter with the 13th letter after it, in the alphabet...
// https://en.wikipedia.org/wiki/ROT13

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	n, e := rr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return n, e
}

// rot13 的转换算法，
func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = b + 13
	case 'a' <= b && b <= 'm':
		b = b + 13
	case 'M' < b && b <= 'Z':
		b = b - 13
	case 'm' < b && b <= 'z':
		b = b - 13
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, r) // You cracked the code!
}
