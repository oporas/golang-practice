package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // close channel (still readable)
	}()

	for n := range c {
		fmt.Println(n)
	}
	//note: main waits for c to close
}
