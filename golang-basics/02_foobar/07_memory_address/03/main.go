package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x) // address in function zero
	fmt.Println(&x)        // address in function zero
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}

// Value is passed to the zero() not the memory address
