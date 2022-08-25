package main

import "fmt"

type structHere struct {
	A, B int
}

func (s *structHere) sum() int {
	return s.A + s.B
}

type interfaceHere interface {
	sum() int
}

func main() {
	var a interfaceHere

	sh := structHere{1, 2}
	os := otherStruct{4, 4}

	a = &sh
	fmt.Println(a.sum())

	a = &os
	fmt.Println(a.sum())

	var b interfaceHere = otherStruct{6, 6}

	fmt.Println(b.sum())
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) sum() int {
	return o.A + o.B
}
