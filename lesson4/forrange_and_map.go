package main

import "fmt"

func main() {
	arr := []int{14, 32, 21, 1, 16}

	for i, j := range arr {
		fmt.Println(i, j)
	}

	firstMap := map[int]string{}
	secondMap := make(map[int]string)

	firstMap[1] = "HELLO"
	firstMap[4] = "WORLD"

	for k, v := range firstMap {
		fmt.Println(k, v)
	}

	value, ok := firstMap[22222]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("key not found")
	}

	fmt.Println(firstMap)
	fmt.Println(secondMap)
}
