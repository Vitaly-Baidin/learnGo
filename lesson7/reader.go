package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")

	arr := make([]byte, 11)

	for {
		s, err := r.Read(arr)
		fmt.Printf("s = %d, err = %v, arr = %v\n", s, err, arr)
		fmt.Printf("arr n bytes %q", arr[:s])
		if err == io.EOF {
			break
		}
	}
}
