package main

import "fmt"

func main() {
	// func

	a := 10
	b := 0

	fmt.Println(sum(a, b))
	fmt.Println(div(a, b))

	//===============================
	//loop

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	sum := 10
	for sum < 1000 {
		sum *= 10
		fmt.Println(sum)
	}

	b = 45

	switch b {
	case 10:
		fmt.Printf("b = %d\n", b)
	case 20:
		fmt.Printf("b = %d\n", b)
	case 45:
		fmt.Printf("b = %d\n", b)
		fallthrough
	case 60:
		fmt.Printf("b = %d\n", b)
	case 70:
		fmt.Printf("b = %d\n", b)
	default:
		fmt.Printf("unknown %d\n", b)
	}

	fmt.Println("==========TEST DEFER===========")
	defer fmt.Println("I'm FIRST DEFER")
	defer fmt.Println("I'm SECOND DEFER")
	fmt.Println("I'm last command")
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("b do not be 0")
	}
	return a / b, nil
}
