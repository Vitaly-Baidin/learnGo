package main

import "fmt"

func main() {
	divide(1, 0)
}

func divide(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(a / b)
}
