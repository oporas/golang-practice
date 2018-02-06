package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1           // blocked at this line since there is no one to read it
	fmt.Println(<-c) // Never gets here
}
