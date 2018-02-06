package main

import (
	"fmt"
)

func main() {
	// c := incrementor()
	// cSum := puller(c)
	// for n := range cSum {
	// 	fmt.Println(n)
	// }
	//No need for the csum
	c := incrementor()
	for n := range puller(c) {
		fmt.Println(n)
	}
}

// optional <- operator specifies the channel direction
func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
