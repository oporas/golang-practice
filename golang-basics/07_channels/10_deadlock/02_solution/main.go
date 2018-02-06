package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // close channel
	}()
	for n := range c { // range through channel output
		fmt.Println(n)
	}
}
