package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("abcdefghijk")

	b := make([]byte, 2)
	for {
		n, err := r.Read(b)
		fmt.Printf("n => %v err => %v b => %v\n", n, err, b)
		if err == io.EOF {
			break
		}
	}
}
