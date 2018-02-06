package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//Run foo&bar and wait before exit from Main
func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i <= 100; i++ {
		fmt.Println("Foo:", i)
		//Add wait for testing...
		time.Sleep(time.Duration(100 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 100; i++ {
		fmt.Println("bar:", i)
		//Add wait for testing...
		time.Sleep(time.Duration(100 * time.Millisecond))
	}
	wg.Done()
}
