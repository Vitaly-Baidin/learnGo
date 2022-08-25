package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go sayHello(ch)

	for i := range ch {
		fmt.Println(i)
	}
}

func say(word string) {
	fmt.Println(word)
}

func sayHello(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		//say("Hello")
		exit <- i
	}

	close(exit)

}
