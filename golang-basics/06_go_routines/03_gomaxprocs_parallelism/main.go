package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	//Automatically added in go1.5~
	runtime.GOMAXPROCS(runtime.NumCPU())
}

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
