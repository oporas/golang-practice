package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)                   // 43
	fmt.Println("a's memory address - ", &a) // 0xc4200140c8

	var b *int = &a
	fmt.Println(b)  // 0xc4200140c8 (memory address)
	fmt.Println(*b) // 43 (deferenced value from b's memory address)

	*b = 42        // change the value at memory address 0xc4200140c8
	fmt.Println(a) // 42
}
