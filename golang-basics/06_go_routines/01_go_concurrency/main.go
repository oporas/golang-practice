package main

import "fmt"

//Prints nothing because main loop gets done before foo/bar has time to run
func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar:", i)
	}
}
