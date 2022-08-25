package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.RWMutex
	c  map[string]int
}

func (c *Counter) CountMe() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func (c *Counter) CountMeAgain() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c
}

func (c *Counter) inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}

func (c *Counter) value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

func main() {
	//key := "test"
	//
	//c := Counter{c: make(map[string]int)}
	//
	//for i := 0; i < 1000; i++ {
	//	go c.inc(key)
	//}
	//
	//time.Sleep(3 * time.Second)
	//
	//fmt.Println(c.value(key))

	//============================================================

	//var wg sync.WaitGroup
	//var counter uint64
	//
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for j := 0; j < 1000; j++ {
	//			atomic.AddUint64(&counter, 1)
	//			counter++
	//		}
	//	}()
	//}
	//
	//wg.Wait()
	//
	//fmt.Println(counter)

	//============================================================

	var once sync.Once
	done := make(chan bool)

	for i := 0; i < 30; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("Hello")
			})
			done <- true
		}()
	}

	for i := 0; i < 30; i++ {
		<-done
	}

}
