package main

import "fmt"

type Point struct { // ~POJO in Java
	X int
	Y int
}

func (p *Point) someMethod() {
	fmt.Println("Call someMethod")
}

func main() {
	a := "Hello"
	b := 42
	fmt.Println(a, b)

	p := a
	a = "World"
	fmt.Println(p, a)
	fmt.Println(&p, &a)
	fmt.Println(&p, &a)

	q := &a
	a = "NEW"
	fmt.Println(*q, a)
}
