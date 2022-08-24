package main

import "fmt"

//======== func callback =======

func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 2)
	fmt.Println(s)
	return result
}

//===============================

//======== Замыкание ============

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(i int) int {
		sum += i
		return sum
	}
}

//===============================

//= Методы структуры + указатели =

type Point struct {
	X, Y int
}

func (p Point) movePoint(x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
}

//===============================

func main() {
	//======== func callback =======
	sumCallback := func(a, b int) int {
		return a + b
	}

	result := doSomething(sumCallback, "sum")
	fmt.Println(result)

	multCallback := func(a, b int) int {
		return a * b
	}

	result = doSomething(multCallback, "mult")
	fmt.Println(result)
	//===============================

	//======== Замыкание ============
	orderPrice := totalPrice(0)
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	//===============================

	//= Методы структуры + указатели =
	p := Point{1, 2}
	fmt.Println(p)
	fmt.Println(p.movePoint(1, 1))
	fmt.Println(p)
	p.movePointPtr(1, 1)
	fmt.Println(p)
	//===============================

}
