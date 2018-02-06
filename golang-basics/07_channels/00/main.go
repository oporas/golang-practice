package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // loop stops until something reads the value of c (in next function)
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	//random sleep for testing, main would otherwise close before funcs
	time.Sleep(time.Second)
}
