package main

import "fmt"

func main() {
	// array
	a := [...]int{1, 2, 3}
	var b [3]int
	b[0] = 4
	b[1] = 5
	b[2] = 6

	fmt.Println(a)
	fmt.Println(b)

	// slice

	c := []int{1, 2, 3}
	c = append(c, 4)

	d := make([]int, 3)
	d[0] = 5
	d[1] = 6
	d[2] = 7

	e := d[1:]

	d[1] = 999

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
