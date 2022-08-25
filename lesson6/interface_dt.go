package main

import "fmt"

func main() {
	var a interface{}
	a = "Hellp"
	fmt.Println(a)
	a = 44
	fmt.Println(a)

	var b interface{} = "some text"

	s := b.(string)
	fmt.Println(s)

	i, ok := b.(int)
	fmt.Println(i, ok)

	var d interface{} = "asdasdasd"

	switch d.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("unknown")
	}
}
