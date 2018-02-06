package main

import "fmt"

func main() {
	c := make(chan int)
	go func() { // run in pararrel to get to the reader line
		c <- 1
	}()
	fmt.Println(<-c)
}
